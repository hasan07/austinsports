package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hasan07/austinsports/lib/log"
	"github.com/hasan07/austinsports/lib/postgres"
)

func (srv *server) UpsertPlayerHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Info(string(body))

	var player postgres.Player
	if err = json.Unmarshal(body, &player); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := srv.DB.UpsertPlayer(r.Context(), player); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, http.StatusOK)
}
