package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

var Mux *http.ServeMux

type Endpoint struct {
	name   string
	method string
	action string
}

type Endpoints struct {
	endpoints []*Endpoint
}

func (e *Endpoints) generateEndpoint(name, action, req string) {
	if IsRequestType(name) {
		e.Bind(&Endpoint{
			name:   strings.ToUpper(name),
			action: action,
		})
	}
}

func (e *Endpoints) Bind(ep *Endpoint) {
	e.endpoints = append(e.endpoints, ep)
}

func (e *Endpoints) BindListeners() {
	for _, end := range e.endpoints {
		Mux.Handle(end.name, end)
	}
}

func (e *Endpoint) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == e.method {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf(`{ "message": %s }`, e.action)))
	}
}

func InitMux() {
	Mux = http.NewServeMux()
}

func NewEndpoint(name, action, method string) *Endpoint {
	return &Endpoint{
		name:   name,
		action: action,
		method: method,
	}
}

func NewEndpoints() *Endpoints {
	return &Endpoints{
		endpoints: []*Endpoint{},
	}
}

func Serve(port int) {
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), Mux))
}

func IsRequestType(name string) bool {
	//TODO
	return true
}

func Delete(w http.ResponseWriter, r *http.Request) {
	//	TODO
}

func Put(w http.ResponseWriter, r *http.Request) {
	//TODO

}

func Head(w http.ResponseWriter, r *http.Request) {
	//TODO

}

func Connect(w http.ResponseWriter, r *http.Request) {
	//TODO
}

func Options(w http.ResponseWriter, r *http.Request) {
	//TODO
}

func Trace(w http.ResponseWriter, r *http.Request) {
	//TODO
}

func Patch(w http.ResponseWriter, r *http.Request) {
	//TODO
}
