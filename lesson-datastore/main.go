package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = "8080"

func main() {
	// TODO - flag for in memory vs other
	s := NewServer(NewMemoryStore())

	addr := fmt.Sprintf(":%s", port)
	log.Printf("running on address: %s", addr)
	http.ListenAndServe(addr, s)
}
