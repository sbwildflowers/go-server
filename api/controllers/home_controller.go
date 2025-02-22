package controllers

import (
	"gotemplate/templates"
	"net/http"
)

func GetHome(res http.ResponseWriter, req *http.Request) {
	component := templates.HomePage()
	cssFiles := []string{"home.css"}
	jsFiles := []string{}
	page := templates.Html(component, cssFiles, jsFiles)
	page.Render(req.Context(), res)
}
