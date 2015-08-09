package palantir

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func init() {
	r := mux.NewRouter()

	http.Handle("/", r)
	r.Handle("/", handle(rootHandler))
}

func rootHandler(w http.ResponseWriter, r *http.Request, c Context) error {
	if r.URL.Path != "/" {
		c.Errorf("Unknow root %s", r.URL.Path)
		http.NotFound(w, r)
		return nil
	}
	fmt.Fprint(w, "Hello hobbits")
	return nil
}
