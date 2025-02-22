package routes

import (
	"gotemplate/controllers"
	"net/http"
)

func SetCallbackPageHandlers(router *http.ServeMux) {
	router.HandleFunc("GET /google/oauth", controllers.HandleGoogleCallback)
}
