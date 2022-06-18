# Lesson Middlewares

This lesson aims to introduce middleware, how to use it and test it.

## How to run

To run in a terminal

```sh
go run main.go
```

## What is middleware?

Middleware is the act of wrapping a http endpoint such that you can log out data, authorise users, protect against panics, record metrics, anything really.

For a good diagram, check out [Dave Stearns blog on middleware](https://drstearns.github.io/tutorials/gomiddleware/).

## How it works

Starting with the http handler, we can make an example Hello World function. After that we'll explore middleware, and after that creating a mux server that uses the middleware. 

```go
func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}
```

After that, we want to make a simple passthrough middleware.
So we define a function that recieves a [http.Handler](https://pkg.go.dev/net/http#Handler) and returns a http.Handler.

> Try writing a middleware that logs before and after the request.

Next we want to define a function that takes in a [*log.Logger](https://pkg.go.dev/log#Logger) and returns the above function.
This requires returning the above function in the log middleware method signature.

After that, try replacing the logger from earlier with the one passed in - in `main.go` we modified the logger so that it prints out the filename as well.

> At this point we should have a function that logs the request data, and a hello world handler func. But they shouldn't be hooked up to anything.

Diving into `main.go`, let's create we want to create a new [mux server](https://pkg.go.dev/net/http#ServeMux).
Initially we want to just [handle](https://pkg.go.dev/net/http#ServeMux.Handle) our endpoints without the middleware.

After this initalise a logger using [log.New](https://pkg.go.dev/log#New) and pass it in to the the logging middleware function, and then wrap the helloworld function.

<details>
  <summary>Click to show solution!</summary>

  ```go
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

	logger.Println("service started on :8080")
	http.ListenAndServe(":8080", mux)
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
  ```
</details>

## Exercises

- Try adding an auth middleware that checks for a username and password. (You can find a solution to this in the `main.go` file)
- Try adding a metrics middleware that records the number of requests made.
- Try adding a panic middleware that catches panics and logs them.
- Try adding tests to one of your middlewares.

## Further Reading

- [Alex Edwards - making and using middleware](https://www.alexedwards.net/blog/making-and-using-middleware)
- [Mat Ryer - Writing middleware in #golang](https://medium.com/@matryer/writing-middleware-in-golang-and-how-go-makes-it-so-much-fun-4375c1246e81)