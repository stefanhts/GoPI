package main

import (
	"GoPI/handlers"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type RequestType int64

var mux *http.ServeMux

const (
	GET RequestType = iota
	POST
	PUT
	DELETE
	HEAD
	CONNECT
	OPTIONS
	TRACE
	PATCH
)

type CRUDType int64

const (
	CREATE CRUDType = iota
	READ
	UPDATE
	DEL
)

type Endpoint struct {
	name        string
	requestType RequestType
	action      string
}

type Endpoints struct {
	endpoints []*Endpoint
}

func (e *Endpoints) generateEndpoint(name, action string, req RequestType) {
	if handlers.IsRequestType(name) {
		e.bind(&Endpoint{
			name:   strings.ToUpper(name),
			action: action,
		})
	}
}

func (e *Endpoints) bind(ep *Endpoint) {
	e.endpoints = append(e.endpoints, ep)
}

func (e *Endpoints) bindListeners() {
	for _, end := range e.endpoints {
		switch end.requestType {
		case GET:
			fmt.Printf("got here")
			mux.HandleFunc(end.name, handlers.Get)
		case POST:
			mux.HandleFunc(end.name, handlers.Post)
		case PUT:
			mux.HandleFunc(end.name, handlers.Put)
		case DELETE:
			mux.HandleFunc(end.name, handlers.Delete)
		case OPTIONS:
			mux.HandleFunc(end.name, handlers.Options)
		case TRACE:
			mux.HandleFunc(end.name, handlers.Trace)
		case PATCH:
			mux.HandleFunc(end.name, handlers.Patch)
		}
		mux.HandleFunc(end.name, end.handleEndpoints)
	}
}

func (e *Endpoint) handleEndpoints(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("ahhhh")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{ "message": e.action }`))
}

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
	mux = http.NewServeMux()
	e := &Endpoint{
		name:        "test",
		requestType: GET,
		action:      "this is a test",
	}
	mux.HandleFunc(e.name, e.handleEndpoints)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
