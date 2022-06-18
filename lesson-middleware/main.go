package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()

	logger := log.New(os.Stdout, "", log.Default().Flags()|log.Lshortfile)
	lm := logMiddleware(logger)

	mux.Handle("/", lm(http.HandlerFunc(helloWorld)))
	mux.Handle("/auth", lm(authMiddleware(http.HandlerFunc(helloAuthorised))))

	logger.Println("service started on :8080")
	http.ListenAndServe(":8080", mux)
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		if username != "username" || password != "hunter2" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func logMiddleware(l *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			l.Println(r.Method, r.URL.Path)
			next.ServeHTTP(w, r)
		})
	}
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func helloAuthorised(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, authorised user!"))
}
