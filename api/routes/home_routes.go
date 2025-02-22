package routes

import (
	"fmt"
	"gotemplate/controllers"
	"net/http"
)

func SetHomePageHandlers(router *http.ServeMux) {
	router.HandleFunc("GET /{$}", controllers.GetHome)
	router.HandleFunc("GET /", http.HandlerFunc(handleNotFound))
}

func handleNotFound(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(res, "404 - Page Not Found")
}
