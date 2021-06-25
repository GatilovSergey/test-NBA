package api

import (
	"encoding/json"
	"net"
	"net/http"
)

func (api *API) NBA(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		if r.RemoteAddr == "" {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("can not find ip address"))
			return
		}
		ip, _, _ := net.SplitHostPort(r.RemoteAddr)

		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(struct {
			Ip string `json:"ip"`
		}{ip})
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}