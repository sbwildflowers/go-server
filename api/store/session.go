package session_store

import (
    "os"
    "github.com/gorilla/sessions"
    "sync"
)

var (
    store *sessions.CookieStore
    once sync.Once
)

func GetStore() *sessions.CookieStore {
    once.Do(func() {
        store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))
    })
    return store
}

