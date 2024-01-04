package model

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/hasan07/austinsports/lib/log"
)

type Options struct {
	Env   string `json:"env"`
	Debug bool   `json:"debug"`
	Port  string `json:"port"`

	SecretFile string `json:"-"`
}

func OptionsFromApp(app *cli.Context) (*Options, error) {
	o := &Options{
		Env:        app.String("env"),
		Debug:      app.Bool("debug"),
		Port:       app.String("port"),
		SecretFile: app.String("secret-file"),
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
	f, err := os.ReadFile(filepath)
	if err != nil {
		return fmt.Errorf("failed to open whisper file: %v", err)
	}
	log.Debugf("got options json: %q", string(f))
	return json.Unmarshal(f, i)
}

func (o *Options) flagOverride(app *cli.Context) {
	if env := app.String("env"); env != "" {
		o.Env = env
	}
}
