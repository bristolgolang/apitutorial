package main

import (
	"log"
	"net/http"
)

type server struct {
	router *http.ServeMux
}

func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func main() {
	s := newServer()
	log.Println("service started on :8080")
	http.ListenAndServe(":8080", s)
}

func newServer() server {
	s := server{router: http.NewServeMux()}
	s.router.HandleFunc("/", helloWorld)
	return s
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}
