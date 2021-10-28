package main

import (
	"embed"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"text/template"
	"time"
)

const port = "8080"

//go:embed static/*
var staticFiles embed.FS

var todoFile = template.Must(template.ParseFS(staticFiles, "static/todoPage.html"))
var quoteFile = template.Must(template.ParseFS(staticFiles, "static/quote.html"))

// TodoPageData is the data structue for the todo page
type TodoPageData struct {
	Title string
	Items []TodoItem
}

// TodoItem is an individual item of the todo list
type TodoItem struct {
	Title string
	Done  bool
}

// Quote is a text attributed to a particular user
type Quote struct {
	Text   string
	Author string
}

func main() {
	http.HandleFunc("/todo", todoPage)
	http.HandleFunc("/quote", quote)
	addr := fmt.Sprintf(":%s", port)
	log.Printf("running on address: %s", addr)
	http.ListenAndServe(addr, nil)
}

func todoPage(w http.ResponseWriter, r *http.Request) {
	data := TodoPageData{
		Title: "My Todo Page",
		Items: []TodoItem{
			{
				Title: "Item 1",
				Done:  false,
			},
			{
				Title: "Item 2",
				Done:  true,
			},
		},
	}

	todoFile.Execute(w, data)
}

func quote(w http.ResponseWriter, r *http.Request) {
	quotes := []Quote{
		{Text: "Simplicity is prerequisite for reliability.", Author: "Edsger W. Dijkstra"},
		{Text: "Good naming is like a good joke. If you have to explain it, it’s not funny. ", Author: "Dave Cheney"},
		{Text: "Don’t comment bad code — rewrite it ", Author: "Brian Kernighan "},
		{Text: "APIs should be easy to use and hard to misuse.", Author: "Josh Bloch"},
		{Text: "Readability is essential for maintainability.", Author: "Mark Reinhold"},
		{Text: "The most important skill for a programmer is the ability to effectively communicate ideas.", Author: "Gastón Jorquera"},
	}

	rand.Seed(time.Now().Unix())
	quote := quotes[rand.Intn(len(quotes))]

	quoteFile.Execute(w, quote)
}
