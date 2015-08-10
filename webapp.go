package palantir

import (
	"appengine/user"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

func init() {
	r := mux.NewRouter()

	http.Handle("/", r)
	r.Handle("/", handle(rootHandler)).Methods("GET")
	r.Handle("/login", handle(loginHandler)).Methods("GET")
	r.Handle("/logout", handle(logoutHandler)).Methods("GET")
}

func rootHandler(w http.ResponseWriter, r *http.Request, c Context) error {
	if r.URL.Path != "/" {
		c.Errorf("Unknow root %s", r.URL.Path)
		http.NotFound(w, r)
		return nil
	}
	tmpl, err := template.ParseFiles("templates/root.tpl", "templates/home.tpl")
	if err != nil {
		return err
	}
	data := struct {
		User *user.User
	}{
		c.user,
	}
	tmpl.Execute(w, data)
	return nil
}

func loginHandler(w http.ResponseWriter, r *http.Request, c Context) error {
	url, err := user.LoginURL(c, "/")
	if err != nil {
		return nil
	}
	http.Redirect(w, r, url, http.StatusFound)
	return nil
}

func logoutHandler(w http.ResponseWriter, r *http.Request, c Context) error {
	url, err := user.LogoutURL(c, "/")
	if err != nil {
		return nil
	}
	http.Redirect(w, r, url, http.StatusFound)
	return nil
}
