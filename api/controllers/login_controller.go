package controllers

import (
	"crypto/rand"
	"encoding/base64"
	"gotemplate/store"
	"gotemplate/templates"
	"net/http"
)

func GetLogin(res http.ResponseWriter, req *http.Request) {
	component := templates.LoginPage()
	cssFiles := []string{}
	jsFiles := []string{}
	page := templates.Html(component, cssFiles, jsFiles)
	page.Render(req.Context(), res)
}

func generateState() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}

func GoogleLogin(res http.ResponseWriter, req *http.Request) {
	config := GetGoogleConfig()
	state := generateState()
	session, _ := session_store.GetStore().Get(req, "session")
	session.Values["oauth-state"] = state
	session.Save(req, res)
	url := config.AuthCodeURL(state)
	http.Redirect(res, req, url, http.StatusTemporaryRedirect)
}
