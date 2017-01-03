package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe("localhost:8001", nil))
}

func hello(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("name")

	if len(name) > 0 {
		fmt.Fprintf(w, "Hello, %s!\n", name)
	} else {
		fmt.Fprintf(w, "Bad request. Parametr name was not specified \n")
		w.WriteHeader(400)
	}
}