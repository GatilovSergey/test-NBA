package api

import (
	"embed"
	"fmt"
	"github.com/GatilovSergey/test-NBA/pkg/nba"
	"github.com/GatilovSergey/test-NBA/pkg/utils"
	"html/template"
	"net/http"
)

// content holds our static web server content.
//go:embed public/views
var content embed.FS

type Todo struct {
	Game []nba.GameInfoPerMin
}

func (api *API) NBA(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		utils.PrettyPrint(Todo{nba.Game})
		t, err := template.ParseFS(content, "public/views/game.html")
		if err!=nil{
			fmt.Println(err)
		}
		t.Execute(w,Todo{nba.Game})


	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (api *API) StartGame(w http.ResponseWriter, r *http.Request) {
	go nba.StartGame()
	http.Redirect(w, r, "/game", 302)
}