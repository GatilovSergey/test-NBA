package server

import (
	"log"
	"net/http"
	"time"

	"github.com/GatilovSergey/test-NBA/pkg/api"
)


type App struct {
	Server   *http.Server
	API      *api.API
}

// NewApp returns initialized struct of main server application.
func NewApp(host string) (*App, error) {
	server := &http.Server{
		Addr:         host,
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	}
	app := &App{
		Server:   server,
	}
	app.API = api.NewApi()
	server.Handler = app.API
	return app, nil
}

func (app *App) Run() error {
	err := app.Server.ListenAndServe()
	if err != nil {
		log.Println("Failed to run server!!!")
		return err
	}
	return nil
}
