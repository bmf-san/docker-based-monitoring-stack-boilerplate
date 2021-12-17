package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		log.Println("OK")
		fmt.Fprintf(w, "Hello World")
	}))
	http.ListenAndServe(":8080", mux)
}
