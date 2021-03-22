package main

import (
	"fmt"
	"net/http"

	"github.com/necronius/go-rest-api/internal/comment"
	"github.com/necronius/go-rest-api/internal/database"
	transportHTTP "github.com/necronius/go-rest-api/internal/transport/http"
)

type App struct{}

func (app *App) Run() error {
	fmt.Println("Settin Up our APP")

	var err error
	db, err := database.NewDatabase()
	if err != nil {
		return err
	}

	commentService := comment.NewService(db)

	handler := transportHTTP.NewHandler(commentService)
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
		return err
	}

	return nil
}

func main() {
	fmt.Println("Go REST api")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up our REST API")
		fmt.Println(err)
	}
}
