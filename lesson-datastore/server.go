package main

import (
	"encoding/json"
	"net/http"
)

// Quote is a thing said by a person
type Quote struct {
	ID     string `json:"id"`
	Text   string `json:"text"`
	Author string `json:"author"`
}

// Datastore is an interface allowing the storage of Quotes
type Datastore interface {
	// insert a quote into the datastore
	Insert(Quote) (Quote, error)
	// list all the quotes in the datastore
	List() ([]Quote, error)
	// pick a random quote from the datastore
	Random() (Quote, error)
}

// server to hold the data store and the routing for the server
type server struct {
	db     Datastore
	router *http.ServeMux
}

// NewServer creates a new server struct, initialised with the routing set
func NewServer(db Datastore) server {
	s := server{
		db:     db,
		router: http.NewServeMux(),
	}
	s.router.HandleFunc("/quotes", s.quotes)
	s.router.HandleFunc("/random", s.randomQuote)
	return s
}

// ServeHTTP has to be implemented for our struct to implement http.Handler
// it simple calls ServerHTTP on its router entity, but this abstraction allows
// us to use our server struct (and with that our datastore) for all incoming requests
func (s server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// quotes handles /quotes endpoint, both POST and GET, so we have to switch on the method
func (s *server) quotes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		// parse the input (no validation is done here)
		var in Quote
		err := json.NewDecoder(r.Body).Decode(&in)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		// insert into the data store
		q, err := s.db.Insert(in)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		// marshals it into bytes to respond with
		b, err := json.Marshal(q)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Write(b)
	case "GET":
		// lists everything in the datastore
		qs, err := s.db.List()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		// marshels it into bytes to respond with
		b, err := json.Marshal(qs)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Write(b)
	}
}

func (s server) randomQuote(w http.ResponseWriter, r *http.Request) {
	// pick a random quote from the datastore
	q, err := s.db.Random()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// marshals it into bytes to respond with
	b, err := json.Marshal(q)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write(b)
}
