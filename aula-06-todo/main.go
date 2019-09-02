package main

import (
	"curso/aula-06-todo/structs"
	"encoding/json"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			data := structs.Response{
				Code: 200,
				Body: "pong",
			}
			json.NewEncoder(w).Encode(data)
		}
	})
	http.ListenAndServe(":3000", mux)
}
