package api


import "net/http"

var _ http.Handler = (*API)(nil)

type API struct {
	mux *http.ServeMux
}

func NewApi() *API {
	api := &API{ mux: http.NewServeMux()}
	api.mux.HandleFunc("/", api.NBA)
	return api
}

func (api *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	api.mux.ServeHTTP(w, r)
}
