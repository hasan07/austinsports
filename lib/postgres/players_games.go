package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/hasan07/austinsports/lib/log"
)

const (
	createPlayersGamesTableStmt = `CREATE TABLE IF NOT EXISTS players_games (
		game_id INTEGER REFERENCES games(id),
		player_id INTEGER REFERENCES players(id),
    	created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    	updated TIMESTAMPTZ NOT NULL DEFAULT NOW(),
		CONSTRAINT players_games_pk PRIMARY KEY(game_id,player_id));`

	upsertPlayersGamesStmt = `INSERT INTO players_games (
		game_id,
		player_id,
		created,
		updated)
	VALUES ($1,$2,NOW(),NOW())
	ON CONFLICT (game_id,player_id)
	DO UPDATE SET
		game_id = $1,
		player_id = $2,
		created = NOW(),
		updated = NOW();`

	getGamesPerPlayerStmt = `
		SELECT
		    g.id,
		    g.name, 
		    g.date,
		    g.location,
		    g.active,
		    g.created,
		    g.updated
		FROM 
		    games g, 
		    players_games pg,
		    players p
		WHERE 
		    g.id=pg.game_id 
		AND 
		    p.email = $1;`
)

type PlayersGames struct {
	GameID   int       `json:"game_id"`
	PlayerID int       `json:"player_id"`
	Created  time.Time `json:"created,omitempty"`
	Updated  time.Time `json:"updated,omitempty"`
}

// UpsertPlayersGames inserts or updates players_games.
func (p *postgresDB) UpsertPlayersGames(ctx context.Context, playersGames PlayersGames) error {
	stmt, err := p.conn.Prepare(upsertPlayersGamesStmt)
	if err != nil {
		return fmt.Errorf("preparing statement %s encountered an error %w", upsertPlayersGamesStmt, err)
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx,
		&playersGames.GameID,
		&playersGames.PlayerID,
	)
	return err
}

func (p *postgresDB) GetGamesPerPlayer(ctx context.Context, playerEmail string) ([]Game, error) {
	rows, err := p.conn.QueryContext(ctx, getGamesPerPlayerStmt, playerEmail)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var games []Game
	for rows.Next() {
		var game Game
		if err := rows.Scan(
			&game.ID,
			&game.Name,
			&game.Date,
			&game.Location,
			&game.Active,
			&game.Created,
			&game.Updated,
		); err != nil {
			log.Error(err)
			return nil, err
		}
		games = append(games, game)
	}

	return games, nil
}
