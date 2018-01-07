package main

import (
	"fmt"
	"net/http"
)

type handle1 struct {
	hello string
}

func main() {
	var h1 handle1
	h1.hello = "aaabbbbccccadddddc"
	http.Handle("/helloworld", &h1)
	fmt.Println(h1.hello)
	http.ListenAndServeTLS(":443", "D:/OnePiece/technology_development/development_environment/214411336860844/214411336860844.pem", "D:/OnePiece/technology_development/development_environment/214411336860844/214411336860844.key", nil)
}
func (h1 *handle1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/helloworld" {
		http.NotFound(w, r)
	}
	fmt.Fprintf(w, h1.hello)
}
