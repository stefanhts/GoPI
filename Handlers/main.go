package handlers

import (
	"net/http"
	"strings"
)

var Requests = []string{"GET", "POST", "PUT", "DELETE", "HEAD", "CONNECT", "OPTIONS", "TRACE", "PATCH"}

func IsRequestType(name string) bool {
	for _, v := range Requests {
		if strings.ToUpper(name) == v {
			return true
		}
	}
	return false
}

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{ "message": "GET request received"}`))
}

func Post(w http.ResponseWriter, r *http.Request) {
	//TODO
}

func Delete(w http.ResponseWriter, r *http.Request) {
	//TODO
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
