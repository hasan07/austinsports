package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/hasan07/austinsports/lib/log"
	"github.com/hasan07/austinsports/lib/postgres"
)

func (srv *server) GetActiveGamesHandler(w http.ResponseWriter, r *http.Request) {

	games, err := srv.DB.GetActiveGames(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(games); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (srv *server) UpsertGameHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Info(string(body))

	var game postgres.Game
	if err = json.Unmarshal(body, &game); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if game.Date.After(time.Now()) {
		game.Active = true
	}

	if err := srv.DB.UpsertGame(r.Context(), game); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, http.StatusOK)
}
