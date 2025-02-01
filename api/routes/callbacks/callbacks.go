package callback_routes

import (
    "gotemplate/controllers/google_oauth"
    "net/http"
)

func SetPageHandlers(router *http.ServeMux) {
    router.HandleFunc("GET /google/oauth", google_oauth_controller.HandleGoogleCallback)
}
