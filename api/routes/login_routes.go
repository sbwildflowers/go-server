package routes

import (
	"gotemplate/controllers"
	"net/http"
)

func SetLoginPageHandlers(router *http.ServeMux) {
	router.HandleFunc("GET /login", controllers.GetLogin)
	router.HandleFunc("GET /login/google", controllers.GoogleLogin)
}
