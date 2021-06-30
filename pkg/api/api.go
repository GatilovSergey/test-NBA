package api

import (
	"fmt"
	"net/http"
)

var _ http.Handler = (*API)(nil)

type API struct {
	*http.ServeMux
}

func NewApi() *API {
	api := &API{  http.NewServeMux()}
	api.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, fmt.Sprintf("<a href=\"%v/start_game\" class=\"button\"> Start Game </a>",r.URL.Host))
	})
	api.HandleFunc("/start_game",api.StartGame)
	api.HandleFunc("/game", api.NBA)
	return api
}

// <p><input type="button" value="Проверить" onclick="isEmail()"></p>
