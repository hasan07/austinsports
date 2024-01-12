package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hasan07/austinsports/lib/log"
	"github.com/hasan07/austinsports/lib/postgres"
)

// UpsertPlayerGameHandler attaches games to a player.
func (srv *server) UpsertPlayerGameHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Info(string(body))

	var playerGames postgres.PlayersGames
	if err = json.Unmarshal(body, &playerGames); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := srv.DB.UpsertPlayersGames(r.Context(), playerGames); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, http.StatusOK)
}

// GetGamesPerPlayerHandler gets all games per player's email.
func (srv *server) GetGamesPerPlayerHandler(w http.ResponseWriter, r *http.Request) {

	games, err := srv.DB.GetGamesPerPlayer(r.Context(), r.URL.Query().Get("player_email"))
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
