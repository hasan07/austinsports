package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/hasan07/austinsports/lib/log"
	"github.com/hasan07/austinsports/lib/model"
)

var (
	createTableStatements = []string{
		createGamesTableStmt,
		createPlayersTableStmt,
		createPlayersGamesTableStmt,
	}
)

type DB interface {
	CreateTables() error

	UpsertGame(ctx context.Context, game Game) error
	GetActiveGames(ctx context.Context) ([]Game, error)

	UpsertPlayer(ctx context.Context, player Player) error

	UpsertPlayersGames(ctx context.Context, playersGames PlayersGames) error
	GetGamesPerPlayer(ctx context.Context, playerEmail string) ([]Game, error)
}

type postgresDB struct {
	opts *model.Options
	conn *sql.DB
	// first map to hold the struct name, second to hold the field name
	modelFields map[string]map[string]bool
}

func New(opts *model.Options) (DB, error) {
	log.Infof("%+v", opts)
	pdb, err := sql.Open("postgres",
		fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
			opts.PostgresUserName, opts.PostgresPassword,
			opts.PostgresHost, opts.PostgresPort, opts.PostgresDB))
	if err != nil {
		return nil, fmt.Errorf("error connecting to blancco DB: %v", err)
	}

	if err := pdb.Ping(); err != nil {
		return nil, err
	}
	return &postgresDB{
		opts: opts,
		conn: pdb,
	}, nil
}

func (p *postgresDB) executeStatement(statement string) error {
	stmt, err := p.conn.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	return err
}

// CreateTables is used to create tables if they do not already exist
func (p *postgresDB) CreateTables() error {
	log.Infof("creating all tables")
	for _, statement := range createTableStatements {
		if err := p.executeStatement(statement); err != nil {
			return err
		}
	}
	return nil
}
