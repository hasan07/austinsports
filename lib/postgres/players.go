package postgres

import (
	"context"
	"fmt"
	"time"
)

const (
	createPlayersTableStmt = `CREATE TABLE IF NOT EXISTS players (
 	ID SERIAL PRIMARY KEY,
 	fname TEXT NOT NULL,
 	lname TEXT NOT NULL,
 	address TEXT NOT NULL,
 	city TEXT NOT NULL,
 	state TEXT NOT NULL,
 	zipcode TEXT NOT NULL, 
 	email TEXT UNIQUE NOT NULL,
 	phone TEXT UNIQUE NOT NULL,
 	created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated TIMESTAMPTZ NOT NULL DEFAULT NOW());`

	upsertPlayerStmt = `INSERT INTO players (
		fname,
		lname,
		address,
		city,
		state,
		zipcode,
		email,
		phone,
		created,
		updated)
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,NOW(),NOW())
	ON CONFLICT (email)
	DO UPDATE SET
		fname = $1,
		lname = $2,
		address = $3,
		city = $4,
		state = $5,
		zipcode = $6,
		email = $7,
		phone = $8,
		created = NOW(),
		updated = NOW();`
)

type Player struct {
	ID      int       `json:"ID"`
	Fname   string    `json:"fname,omitempty"`
	Lname   string    `json:"lname,omitempty"`
	Address string    `json:"address,omitempty"`
	City    string    `json:"city,omitempty"`
	State   string    `json:"state,omitempty"`
	Zipcode string    `json:"zipcode,omitempty"`
	Email   string    `json:"email,omitempty"`
	Phone   string    `json:"phone,omitempty"`
	Created time.Time `json:"created,omitempty"`
	Updated time.Time `json:"updated,omitempty"`
}

// UpsertPlayer inserts or updates players.
func (p *postgresDB) UpsertPlayer(ctx context.Context, player Player) error {
	stmt, err := p.conn.Prepare(upsertPlayerStmt)
	if err != nil {
		return fmt.Errorf("preparing statement %s encountered an error %w", upsertPlayerStmt, err)
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx,
		&player.Fname,
		&player.Lname,
		&player.Address,
		&player.City,
		&player.State,
		&player.Zipcode,
		&player.Email,
		&player.Phone,
	)
	return err
}
