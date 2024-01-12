package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (srv *server) v1Router(r *mux.Router) {
	v1 := r.PathPrefix("/api/v1").Subrouter()

	// Games
	gamesRouter := v1.PathPrefix("/games").Subrouter()
	gamesRouter.HandleFunc("", srv.GetActiveGamesHandler).Methods(http.MethodGet)
	gamesRouter.HandleFunc("/create", srv.UpsertGameHandler).Methods(http.MethodPost)

	// Players
	playersRouter := v1.PathPrefix("/players").Subrouter()
	playersRouter.HandleFunc("", srv.GetActiveGamesHandler).Methods(http.MethodGet)
	playersRouter.HandleFunc("/create", srv.UpsertPlayerHandler).Methods(http.MethodPost)

	// PlayersGames junction
	playerGameRouter := v1.PathPrefix("/player_games").Subrouter()
	playerGameRouter.HandleFunc("", srv.GetGamesPerPlayerHandler).Methods(http.MethodGet)
	playerGameRouter.HandleFunc("/create", srv.UpsertPlayerGameHandler).Methods(http.MethodPost)

}
