package login_routes

import (
	"metrics/controllers/login"
	"net/http"
)

func SetPageHandlers(router *http.ServeMux) {
    router.HandleFunc("GET /login", login_controller.GetLogin)
	router.HandleFunc("GET /login/google", login_controller.GoogleLogin)
}
