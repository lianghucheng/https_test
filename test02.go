package main

import (
	"fmt"
	"net/http"
)

type handletest struct {
}

func main() {
	var server handletest
	http.Handle("/hhh", &server)
	http.ListenAndServe("localhost:8080", nil)
}

func (server *handletest) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hhh" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "aaaaa")
}

/*package main

import (
	"fmt"
	"net/http"
)

type handle struct {
}

func main() {
	var server handle
	http.Handle("/hhh", &server)
	http.ListenAndServe("localhist:8080", nil)
}

func (h *handle) hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hhh" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "aaaa")
}
*/
