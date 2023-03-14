package main

import (
	"fmt"
	"net/http"
)

func tempServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}

func main() {
	http.HandleFunc("/", tempServer)
	http.ListenAndServe("", nil)
}
