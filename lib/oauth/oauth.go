package oauth

import (
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"

	"github.com/hasan07/austinsports/lib/log"
	"github.com/hasan07/austinsports/lib/model"
)

const (
	MaxAge = 86400 * 30
)

func New(opts *model.Options) {
	log.Info("HERE!!!!!")
	store := sessions.NewCookieStore([]byte("SECRET_SESSION_KEY"))
	store.MaxAge(MaxAge)

	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = opts.Env == "prod"

	gothic.Store = store

	goth.UseProviders(
		google.New(opts.GoogleID, opts.GoogleKey, "http://127.0.0.1:8080/auth/google/callback", "email", "profile"),
	)
}
