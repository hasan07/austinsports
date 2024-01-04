package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (srv *server) v1Router(r *mux.Router) {
	v1 := r.PathPrefix("/api/v1").Subrouter()

	// Games
	gamesRouter := v1.PathPrefix("/games").Subrouter()
	gamesRouter.HandleFunc("/get", srv.GamesHandler).Methods(http.MethodGet)

}

func (srv *server) GamesHandler(w http.ResponseWriter, r *http.Request) {
}
