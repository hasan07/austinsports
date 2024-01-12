package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/markbates/goth/gothic"

	"github.com/hasan07/austinsports/lib/log"
)

func (srv *server) v1Router(r *mux.Router) {
	v1 := r.PathPrefix("/api/v1").Subrouter()
	authRouter := r.PathPrefix("").Subrouter()
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

	// Auth
	authRouter.HandleFunc("/auth/{provider}/callback", srv.getAuthcallBackFunc)
	authRouter.HandleFunc("/auth/{provider}", func(res http.ResponseWriter, req *http.Request) {
		gothic.BeginAuthHandler(res, req)
	})

}

func (srv *server) getAuthcallBackFunc(w http.ResponseWriter, r *http.Request) {
	provider := mux.Vars(r)["provider"]
	fmt.Println(provider)
	r = r.WithContext(context.WithValue(r.Context(), "provider", provider))
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Infof("%+v", user)

	http.Redirect(w, r, "http://127.0.0.1:5173", http.StatusFound)
}
