package main

import (
	"log"
	"net/http"
)

type Server struct{}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{ "message": "GET request received"}`))
	case "POST":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{ "message": "POST request received"}`))
	case "PUT":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{ "message": "PUT request received"}`))
	case "DELETE":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{ "message": "DELETE request received"}`))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{ "message": "not found"}`))

	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":5000", mux))
}
