package api

import (
	"fmt"
	"net/http"

	"github.com/NYTimes/gziphandler"
	"github.com/gorilla/mux"
	"github.com/urfave/cli/v2"

	"github.com/hasan07/austinsports/lib/log"
	"github.com/hasan07/austinsports/lib/model"
)

var Cmd = &cli.Command{
	Name:   "api",
	Action: run,
}

func run(app *cli.Context) error {
	opts, err := model.OptionsFromApp(app)
	if err != nil {
		return err
	}
	log.Info(opts)

	srv, err := New(opts)
	if err != nil {
		return fmt.Errorf("failed to create server: %v", err)
	}

	return srv.Serve()

	// TODO(hmachlab): Implement server logic.
	return nil
}

type server struct {
	opts *model.Options
}

func New(opts *model.Options) (*server, error) {
	return &server{opts: opts}, nil
}

// Serve starts the http listener.
func (srv *server) Serve() error {
	r := mux.NewRouter()
	srv.v1Router(r)
	log.Infof("api started; listening on port %s", srv.opts.Port)
	return http.ListenAndServe(fmt.Sprintf(":%s", srv.opts.Port), gziphandler.GzipHandler(r))
}
