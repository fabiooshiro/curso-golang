package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/go/lalala", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ola mundo go lalala\n\n"))
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ola mundo\n\n"))
	})
	mux.HandleFunc("/go/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ola mundo go\n\n"))
	})

	http.ListenAndServe("localhost:3000", mux)
}
