package middleware

import (
	"context"
	"metrics/controllers/google_oauth"
	"metrics/utils"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func VerifyUser(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        unprotected_routes := []string{"/login", "/login/google", "/google/oauth"}
        if array_utils.ArrayContains(unprotected_routes, r.URL.Path) {
            next.ServeHTTP(w, r)
        } else {
            cookie, err := r.Cookie("auth_token")
            if err != nil {
                http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
                return
            }

            token, err := jwt.ParseWithClaims(cookie.Value, &google_oauth_controller.Claims{}, func(token *jwt.Token) (interface{}, error) {
                return []byte(os.Getenv("JWT_SECRET")), nil
            })

            if err != nil || !token.Valid {
                http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
                return
            }

            if claims, ok := token.Claims.(*google_oauth_controller.Claims); ok {
                ctx := context.WithValue(r.Context(), "user", claims)
                next.ServeHTTP(w, r.WithContext(ctx))
            } else {
                http.Error(w, "Unauthorized", http.StatusUnauthorized)
                return
            }
        }
    })
}
