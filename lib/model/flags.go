package model

import "github.com/urfave/cli/v2"

var MainFlags = []cli.Flag{
	&cli.BoolFlag{
		Name:    "debug",
		Usage:   "enable debug log level",
		EnvVars: []string{"DEBUG"},
	},
}

var DefaultAPIFlags = []cli.Flag{
	&cli.StringFlag{
		Name:    "env",
		Usage:   "env",
		Value:   "dev",
		EnvVars: []string{"ENV"},
	},
	&cli.StringFlag{
		Name:    "port",
		Usage:   "server port",
		EnvVars: []string{"API_PORT"},
		Value:   "8080",
	},
}

var DefaultDBFlags = []cli.Flag{
	&cli.StringFlag{
		Name:  "postgres-pw",
		Usage: "Postgres Password",
	},
	&cli.StringFlag{
		Name:  "postgres-un",
		Usage: "Postgres Username",
	},
	&cli.StringFlag{
		Name:  "postgres-host",
		Usage: "Postgres host",
	},
	&cli.StringFlag{
		Name:  "postgres-port",
		Usage: "Postgres port",
		Value: "5432",
	},
	&cli.StringFlag{
		Name:  "postgres-db",
		Usage: "Postgres DB",
		Value: "austinsports",
	},
}

var SecretFlags = []cli.Flag{
	&cli.StringFlag{
		Name:    "secret-file, sf",
		Usage:   "file path of a json file which is suitable for supplying options for the api",
		EnvVars: []string{"SECRET_FILE"},
	},
}

func JoinFlags(flags ...[]cli.Flag) []cli.Flag {
	var out []cli.Flag
	for _, f := range flags {
		out = append(out, f...)
	}
	return out
}
