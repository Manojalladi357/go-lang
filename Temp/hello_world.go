package main

import (
	"fmt"
	"net/http"
)

func serverPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world!")
}

func main() {

	http.HandleFunc("/", serverPage)
	http.ListenAndServe("", nil)
}
