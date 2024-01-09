package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/hasan07/austinsports/lib/log"
)

const (
	createGamesTableStmt = `CREATE TABLE IF NOT EXISTS games (
 	ID SERIAL PRIMARY KEY,
 	date TIMESTAMPTZ NOT NULL,
 	location TEXT NOT NULL,
 	name TEXT NOT NULL,
 	description TEXT NOT NULL,
 	address TEXT NOT NULL,
 	city TEXT NOT NULL,
 	state TEXT NOT NULL,
 	zipcode TEXT NOT NULL,
 	price NUMERIC(7,5),
 	active BOOLEAN,
 	created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated TIMESTAMPTZ NOT NULL DEFAULT NOW());`

	upsertGameStmt = `INSERT INTO games (
		date,
		location,
		name,
		description,
		address,
		city,
		state,
		zipcode,
        price,           
        active,
		created,
		updated)
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,NOW(),NOW())
	ON CONFLICT (ID)
	DO UPDATE SET
		date = $1,
		location = $2,
		name = $3,
		description = $4,
		address = $5,
		city = $6,
		state = $7,
		zipcode = $8,
	    price = $9,
	    active = $10,
		created = NOW(),
		updated = NOW();`

	getActiveGamesStmt = `
	SELECT 
		ID,
		date,
		location,
		name,
		description,
		address,
		city,
		state,
		zipcode,
	    price,
		active,
		created,
		updated
	FROM 
	    games
	WHERE active = true;

`
)

type Game struct {
	ID          int       `json:"ID,omitempty"`
	Date        time.Time `json:"date,omitempty"`
	Location    string    `json:"location,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	Address     string    `json:"address,omitempty"`
	City        string    `json:"city,omitempty"`
	State       string    `json:"state,omitempty"`
	Zipcode     string    `json:"zipcode,omitempty"`
	Price       float64   `json:"price,omitempty"`
	Active      bool      `json:"active"`
	Created     time.Time `json:"created,omitempty"`
	Updated     time.Time `json:"updated,omitempty"`
}

// UpsertGame inserts or updates game.
func (p *postgresDB) UpsertGame(ctx context.Context, game Game) error {
	stmt, err := p.conn.Prepare(upsertGameStmt)
	if err != nil {
		return fmt.Errorf("preparing statement %s encountered an error %w", upsertGameStmt, err)
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx,
		&game.Date,
		&game.Location,
		&game.Name,
		&game.Description,
		&game.Address,
		&game.City,
		&game.State,
		&game.Zipcode,
		&game.Price,
		&game.Active,
	)
	return err
}

func (p *postgresDB) GetActiveGames(ctx context.Context) ([]Game, error) {
	rows, err := p.conn.QueryContext(ctx, getActiveGamesStmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var games []Game
	for rows.Next() {
		var game Game
		if err := rows.Scan(
			&game.ID,
			&game.Date,
			&game.Location,
			&game.Name,
			&game.Description,
			&game.Address,
			&game.City,
			&game.State,
			&game.Zipcode,
			&game.Price,
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
