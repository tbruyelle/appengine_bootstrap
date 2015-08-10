package appengine_bootstrap

import (
	"net/http"

	"appengine/user"

	"appengine"
)

type Context struct {
	appengine.Context
	user *user.User
}

var (
	contexter = func(r *http.Request) appengine.Context {
		return appengine.NewContext(r)
	}
)

type ctxHandler func(http.ResponseWriter, *http.Request, Context) error

func handle(h ctxHandler) http.HandlerFunc {
	return _handle(h, false)
}

func handleLogged(h ctxHandler) http.HandlerFunc {
	return _handle(h, true)
}

func _handle(h ctxHandler, assertLogged bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := Context{}
		ctx.Context = contexter(r)
		ctx.user = user.Current(ctx)
		if assertLogged && ctx.user == nil {
			// No logged user, redirect to root
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		err := h(w, r, ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
