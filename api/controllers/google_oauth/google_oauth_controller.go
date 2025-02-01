package google_oauth_controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

    "gotemplate/store"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleUser struct {
    Email   string `json:"email"`
    VerifiedEmail bool `json:"verified_email"`
    Name    string `json:"name"`
    Picture string `json:"picture"`
}

type Claims struct {
    Email string `json:"email"`
    Name string `json:"name"`
    Picture string `json:"picture"`
    jwt.RegisteredClaims
}

func HandleGoogleCallback(res http.ResponseWriter, req *http.Request) {
    session, _ := session_store.GetStore().Get(req, "session")
    expectedState, ok := session.Values["oauth-state"].(string)
    if !ok {
        http.Error(res, "Invalid session", http.StatusBadRequest)
        return
    }

    if req.FormValue("state") != expectedState {
        http.Error(res, "state is invalid", http.StatusBadRequest)
        return
    }

    delete(session.Values, "oauth-state")
    session.Save(req, res)

    code := req.FormValue("code")
    token, err := config.Exchange(req.Context(), code)
    if err != nil {
        fmt.Println(err)
        http.Error(res, "Failed to exchange token", http.StatusInternalServerError)
        return
    }

    client := config.Client(req.Context(), token)
    user_info, err := getUserInfo(client)
    if err != nil {
        fmt.Println(err)
        http.Error(res, "Failed to get user info", http.StatusInternalServerError)
        return
    }

    jwt_token, err := createToken(user_info)
    if err != nil {
        http.Error(res, "Failed to create token", http.StatusInternalServerError)
    }

    setCookie(res, jwt_token)

    http.Redirect(res, req, "/", http.StatusTemporaryRedirect)
}



var config *oauth2.Config

func InitConfig() {
    config = &oauth2.Config{
        ClientID: os.Getenv("GOOGLE_CLIENT_ID"),
        ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
        RedirectURL: os.Getenv("CALLBACK_URL"),
        Scopes: []string{
            "https://www.googleapis.com/auth/userinfo.email",
            "https://www.googleapis.com/auth/userinfo.profile",
        },
        Endpoint: google.Endpoint,
    }
}

func GetGoogleConfig() *oauth2.Config {
    return config
}

func getUserInfo(client *http.Client) (*GoogleUser, error) {
    res, err := client.Get("https://www.googleapis.com/userinfo/v2/me")
    if err != nil {
        return nil, err
    }
    defer res.Body.Close()

    body, err := io.ReadAll(res.Body)
    if err != nil {
        return nil, fmt.Errorf("failed to read response body: %v", err)
    }
    if res.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("bad response: %d, body: %s", res.StatusCode, string(body))
    }

    var user GoogleUser
    if err := json.Unmarshal(body, &user); err != nil {
        return nil, fmt.Errorf("failed to decode response: %v, body: %s", err, string(body))
    }

    return &user, nil
}

func createToken(user *GoogleUser) (string, error) {
    claims := Claims{
        Email: user.Email,
        Name: user.Name,
        Picture: user.Picture,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
            IssuedAt: jwt.NewNumericDate(time.Now()),
            NotBefore: jwt.NewNumericDate(time.Now()),
            Issuer: "gotemplate",
            Subject: user.Email,
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    signed_token, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
    if err != nil {
        return "", err
    }

    return signed_token, nil
}

func setCookie(res http.ResponseWriter, token string) {
    var secure bool
    if os.Getenv("ENVIRON") == "prod" {
        secure = true
    } else {
        secure = false
    }
    http.SetCookie(res, &http.Cookie{
        Name: "auth_token",
        Value: token,
        Path: "/",
        HttpOnly: true,
        Secure: secure,
        SameSite: http.SameSiteLaxMode,
        MaxAge: 24 * 60 * 60,
    })
}
