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
}

func rootHandler(w http.ResponseWriter, r *http.Request, c Context) error {
	if r.URL.Path != "/" {
		c.Errorf("Unknow root %s", r.URL.Path)
		http.NotFound(w, r)
		return nil
	}
	tmpl, err := template.ParseFiles("static/root.html", "static/home.html")
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
