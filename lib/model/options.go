package model

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/hasan07/austinsports/lib/log"

	_ "github.com/lib/pq"
)

type Options struct {
	Env   string `json:"env"`
	Debug bool   `json:"debug"`
	Port  string `json:"port"`

	SecretFile string `json:"-"`

	// Postgres
	PostgresUserName string `json:"postgres_user_name"`
	PostgresPassword string `json:"postgres_password"`
	PostgresPort     int    `json:"postgres_port"`
	PostgresHost     string `json:"postgres_host"`
	PostgresDB       string `json:"postgres_db"`

	GoogleID  string `json:"google_id"`
	GoogleKey string `json:"google_key"`
}

func OptionsFromApp(app *cli.Context) (*Options, error) {
	o := &Options{
		Env:              app.String("env"),
		Debug:            app.Bool("debug"),
		Port:             app.String("port"),
		SecretFile:       app.String("secret-file"),
		PostgresUserName: app.String("pg-username"),
		PostgresPassword: app.String("pg-password"),
		PostgresPort:     app.Int("pg-port"),
		PostgresHost:     app.String("pg-host"),
		PostgresDB:       app.String("pg-db"),
		GoogleID:         app.String("google-id"),
		GoogleKey:        app.String("google-key"),
	}

	if o.SecretFile != "" {
		err := SecretFromFile(o.SecretFile, o)
		o.flagOverride(app)
		return o, err
	}
	return o, nil
}

// SecretFromFile reads json file into the given interface.
func SecretFromFile(filepath string, i interface{}) error {
	log.Info("getting secrets from file")
	f, err := os.ReadFile(filepath)
	if err != nil {
		return fmt.Errorf("failed to open secrets file: %v", err)
	}
	log.Debugf("got options json: %q", string(f))
	return json.Unmarshal(f, i)
}

func (o *Options) flagOverride(app *cli.Context) {
	if env := app.String("env"); env != "" {
		o.Env = env
	}
}
