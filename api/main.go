package main

import (
	"fmt"
	"log"
	"metrics/controllers/google_oauth"
    "metrics/database"
	"metrics/middleware"
	"metrics/routes/callbacks"
	"metrics/routes/home"
	"metrics/routes/login"
	"net/http"
    "os"
	"github.com/joho/godotenv"
)

func init() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    google_oauth_controller.InitConfig()

    err = database.ConnectToDatabase() 
    if err != nil {
        log.Fatal("Connection to DB failed")
    }
}

func main() {
    router := http.NewServeMux()

    // handle static resources
    staticFs := http.FileServer(http.Dir("./static"))
    router.Handle("GET /static/", http.StripPrefix("/static", staticFs))

    // routes
    pageRouter := http.NewServeMux()
    home_routes.SetPageHandlers(pageRouter)
    login_routes.SetPageHandlers(pageRouter)
    callback_routes.SetPageHandlers(pageRouter)

    stack := middleware.CreateMiddlewareStack(
        middleware.VerifyUser,
        middleware.Logging,
    )
    router.Handle("/", stack(pageRouter))

    // serve
    port := os.Getenv("PORT")
    server := http.Server{
        Addr: fmt.Sprintf(":%s", port),
        Handler: router, 
    }
    fmt.Println("Server listening on port", port)
    server.ListenAndServe()
}
