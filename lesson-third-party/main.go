package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

var client *http.Client

const host = "https://quote-garden.herokuapp.com/quotes"

// Quotes is the response type from a search
type Quotes struct {
	Count   int     `json:"count"`
	Results []Quote `json:"results"`
}

// Quote is a quote
type Quote struct {
	// ID          string `json:"_id"`
	QuoteText   string `json:"quoteText"`
	QuoteAuthor string `json:"quoteAuthor"`
}

func (q *Quote) String() string {
	return fmt.Sprintf(`%s
> %s
`, q.QuoteText, q.QuoteAuthor)
}

func main() {
	client = http.DefaultClient
	client.Timeout = 10 * time.Second // set a timeout limit of 10 seconds
	http.HandleFunc("/quotes/random", getRandomQuote)
	http.HandleFunc("/quotes/search/", getSearchQuote)
	http.ListenAndServe(":8080", nil)
}

// getRandomQuote gets a random quote from our source
func getRandomQuote(w http.ResponseWriter, r *http.Request) {
	resp, err := client.Get(fmt.Sprintf("%s/random", host))
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Errorf("getting response: %v", err.Error()).Error()))
		return
	}
	defer resp.Body.Close()
	// io.Copy(w, resp.Body)
	var q Quote
	err = json.NewDecoder(resp.Body).Decode(&q)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Errorf("decoding response: %v", err.Error()).Error()))
		return
	}
	w.Write([]byte(q.String()))
}

// getSearchQuote returns a list of quotes
func getSearchQuote(w http.ResponseWriter, r *http.Request) {
	searchTerm := strings.TrimPrefix(r.URL.Path, "/quotes/search/")
	resp, err := client.Get(fmt.Sprintf("%s/search/%s", host, searchTerm))
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Errorf("getting response: %v", err.Error()).Error()))
		return
	}
	defer resp.Body.Close()
	// io.Copy(w, resp.Body)
	var quotes Quotes
	err = json.NewDecoder(resp.Body).Decode(&quotes)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Errorf("decoding response: %v", err.Error()).Error()))
		return
	}
	i := 0
	for _, q := range quotes.Results {
		w.Write([]byte(q.String()))
		i++
		if i > 4 {
			break
		}
	}
}
