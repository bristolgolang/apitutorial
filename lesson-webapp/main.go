package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"
)

const port = "8080"

var todoTmpl = template.Must(template.ParseFiles("lesson-webapp/static/todoPage.html"))
var quoteTmpl = template.Must(template.ParseFiles("lesson-webapp/static/quote.html"))

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
			TodoItem{
				Title: "Item 1",
				Done:  false,
			},
			TodoItem{
				Title: "Item 2",
				Done:  true,
			},
		},
	}
	todoTmpl.Execute(w, data)
}

func quote(w http.ResponseWriter, r *http.Request) {
	quotes := []Quote{
		Quote{Text: "Simplicity is prerequisite for reliability.", Author: "Edsger W. Dijkstra"},
		Quote{Text: "Good naming is like a good joke. If you have to explain it, it’s not funny. ", Author: "Dave Cheney"},
		Quote{Text: "Don’t comment bad code — rewrite it ", Author: "Brian Kernighan "},
		Quote{Text: "APIs should be easy to use and hard to misuse.", Author: "Josh Bloch"},
		Quote{Text: "Readability is essential for maintainability.", Author: "Mark Reinhold"},
		Quote{Text: "The most important skill for a programmer is the ability to effectively communicate ideas.", Author: "Gastón Jorquera"},
	}

	rand.Seed(time.Now().Unix())
	quote := quotes[rand.Intn(len(quotes))]

	quoteTmpl.Execute(w, quote)
}
