package sessionstore

import (
	"myapp/internal/config"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte(config.SessionSecret))

func GetStore() *sessions.CookieStore {
	store.Options = &sessions.Options{
		HttpOnly: true,
		MaxAge:   60 * 60 * 24 * 7,
	}

	return store
}
