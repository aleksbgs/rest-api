package main

import (
	"fmt"
	transportHTTP "github.com/aleksbgs/rest-api/internal/transport/http"
	"net/http"
)

// App- the struct which contains things like pointers
// to database connection
type App struct{}

// Run- sets up application
func (app *App) Run() error {
	fmt.Println("Settings Up Our APP")
	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
		return err
	}
	return nil
}

func main() {
	fmt.Println("Go rest APi Course")

	app := App{}

	if err := app.Run(); err != nil {
		fmt.Println("Error starting up our REST API")
		fmt.Println(err)
	}
}
