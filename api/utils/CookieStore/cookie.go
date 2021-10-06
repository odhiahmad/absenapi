package CookieStore

import "github.com/gorilla/sessions"

type AppCookieStore struct {
	Store *sessions.CookieStore
}

func NewAppCookieStore() *AppCookieStore {
	var store = sessions.NewCookieStore([]byte("BRI..."))
	store.MaxAge(60)
	return &AppCookieStore{store}
}
