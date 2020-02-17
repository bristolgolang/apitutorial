package main

import (
	"fmt"
)

// memory is an in memory implementation of Datastore
type memory struct {
	quotes map[string]Quote
}

// NewMemoryStore creates a new in memory data store
func NewMemoryStore() *memory {
	return &memory{
		quotes: make(map[string]Quote),
	}
}

// Insert implements the interface
func (m *memory) Insert(q Quote) (Quote, error) {
	if q.ID != "" {
		return Quote{}, fmt.Errorf("ID of quote should not be set: %s", q.ID)
	}
	id := randomID(8)

	q.ID = id
	m.quotes[q.ID] = q

	return q, nil
}

// Random implements the datastore interface
func (m *memory) Random() (Quote, error) {
	// two options on picking a random item from a map in Go
	// either iterate the map and break immediately (maps aren't guaranteed to be deterministic in ordering)
	// or get the keys into a slice and choose an entry at random
	var k string
	var q Quote
	for k, q = range m.quotes {
		break
	}
	if k == "" {
		return Quote{}, fmt.Errorf("no quotes in data store")
	}
	return q, nil
}

// List implements the datastore interface
func (m *memory) List() ([]Quote, error) {
	quotes := make([]Quote, len(m.quotes))
	for _, v := range m.quotes {
		quotes = append(quotes, v)
	}
	return quotes, nil
}
