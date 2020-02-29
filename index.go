package goserver

import (
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	cmd := r.FormValue("cmd")
	fmt.Fprint(w, cmd)
}
func init() {
	http.HandleFunc("/", IndexHandler)
}
