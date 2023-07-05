package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/arturfil/yt-go-coffee-server/db"
	"github.com/arturfil/yt-go-coffee-server/router"
	"github.com/arturfil/yt-go-coffee-server/services"
	"github.com/joho/godotenv"
)

type Config struct {
    Port string
}

type Application struct {
    Config Config
    Models services.Models 
}

func (app *Application) Serve() error {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error leading .env file")
    }
    port := os.Getenv("PORT")
    fmt.Println("API is listening on port", port)

    srv := &http.Server {
        Addr: fmt.Sprintf(":%s", port),
        Handler: router.Routes(), 
    }
    return srv.ListenAndServe()
}

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error leading .env file")
    }

    cfg := Config {
        Port: os.Getenv("PORT"),
    }

    dsn := os.Getenv("DSN")
    dbConn, err := db.ConnectPostgres(dsn)
    if err != nil {
        log.Fatal("Cannot connect to database")
    }

    defer dbConn.DB.Close()

    app := &Application {
        Config: cfg,
        Models: services.New(dbConn.DB),
    }

    err = app.Serve()
    if err != nil {
        log.Fatal(err)
    }
}
