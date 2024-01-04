package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/hasan07/austinsports/cmds/api"
	"github.com/hasan07/austinsports/lib/log"
	"github.com/hasan07/austinsports/lib/model"
)

var (
	Version = "dev"

	// WantArgsLen expected number of arguments.
	WantArgsLen = 1

	// ErrInvalidNumberOfArguments returned when using an invalid number of arguments.
	ErrInvalidNumberOfArguments = errors.New("invalid number of arguments")
)

func cmd(app *cli.Context) error {
	if app.NArg() != WantArgsLen {
		return fmt.Errorf("%w: %d != %d", ErrInvalidNumberOfArguments, app.NArg(), WantArgsLen)
	}
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "austin-sports"
	app.Usage = "austin-sports"
	app.Version = Version
	app.EnableBashCompletion = true
	app.Flags = model.MainFlags
	app.Before = func(ctx *cli.Context) error {
		return nil
	}
	app.After = func(ctx *cli.Context) error {
		c := log.GlobalConfig()
		c.DisableStacktrace = true
		if err := log.SetGlobalConfig(c); err != nil {
			return err
		}
		return nil
	}

	app.Action = cmd
	app.Commands = []*cli.Command{
		api.Cmd,
	}

	if err := app.Run(os.Args); err != nil {
		log.Error(err)
	}
}
