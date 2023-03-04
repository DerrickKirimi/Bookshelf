package app

import (
	"net/http"
)

func (app *App) HandleIndex(W http.ResponseWriter, _ *http.Request) {
	W.Header().Set("Content-Length", "12")
	W.Header().Set("Content-Type", "text/plain; charset=utf-8")
	W.Header().Set("X-Content-Type-Options", "nosniff")
	W.Write([]byte("Welcome Home"))
}
