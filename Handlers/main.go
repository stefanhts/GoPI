package handlers

import (
	"net/http"
	"strings"
	"fmt"
)

//TODO: move all Endpoints structs and such to this module

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
	path := r.URL.Path
	w.Write([]byte(fmt.Sprintf(`{ "message": %s}`, path)))
	
}

//func (e *Endpoints) GetRecord(path string){
//	//TODO: implement this
//}

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
