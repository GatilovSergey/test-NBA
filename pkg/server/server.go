package server

import (
	"embed"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/GatilovSergey/test-NBA/pkg/api"

)

// content holds our static web server content.
//go:embed public/views/*.html
var content embed.FS

type App struct {
	Template *template.Template
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
	t := template.Must(template.ParseFS(content, "public/views/*.html"))
	app := &App{
		Template: t,
		Server:   server,
		//Srv:      lobby.NewLobby(),
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
