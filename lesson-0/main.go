package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloWorld)
	log.Println("service started on :8080")
	http.ListenAndServe(":8080", nil)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}
