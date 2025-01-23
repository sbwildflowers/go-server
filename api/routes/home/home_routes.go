package home_routes

import (
	"fmt"
	"metrics/controllers/home"
	"net/http"
)

func SetPageHandlers(router *http.ServeMux) {
    router.HandleFunc("GET /{$}", home_controller.GetHome)
    router.HandleFunc("GET /", http.HandlerFunc(handleNotFound))
}

func handleNotFound(res http.ResponseWriter, req *http.Request) {
    res.WriteHeader(http.StatusNotFound)
    fmt.Fprintf(res, "404 - Page Not Found")
}
