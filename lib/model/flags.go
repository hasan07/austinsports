package model

import "github.com/urfave/cli/v2"

var MainFlags = []cli.Flag{
	&cli.StringFlag{
		Name:    "env",
		Usage:   "env",
		Value:   "dev",
		EnvVars: []string{"ENV"},
	},
	&cli.StringFlag{
		Name:    "debug",
		Usage:   "enable debug log level",
		EnvVars: []string{"DEBUG"},
	},
	&cli.StringFlag{
		Name:    "port",
		Usage:   "server port",
		EnvVars: []string{"API_PORT"},
		Value:   "8080",
	},
}

var DefaultWhisperFlags = []cli.Flag{
	&cli.StringFlag{
		Name:    "secret-file, sf",
		Usage:   "file path of a json file which is suitable for supplying options for the api",
		EnvVars: []string{"SECRET_FILE"},
	},
}
