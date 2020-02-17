# Lesson Datastore


## How to run

From the root directory of this repository;

```sh
go run lesson-datastore/*.go
```

This command is a little different to run, as we need Go to read all the files within this directory, we've split this into multiple files for simplicity.

```sh
curl -X POST --data '{"text": "Alex is awesome!", "author": "AJT"}' localhost:8080/quotes
curl localhost:8080/random
```

This would post a quote of `Alex is awesome!` by `AJT` into our system, we'd then retrieve a random item from our datastore (which is likely to be our text we just put in).

## How it works

### The main file

There are a few major components here as we begin to explore the power of Go's interfaces.
The first thing we should explore is the `main.go` file as this is where our application kicks off.
This is a simple file that just calls `NewServer`, with the arguments from a `NewMemoryStore`.
After that we create an address, log out where our server is running.
Finally we call [`http.ListenAndServe(addr, s)`](https://golang.org/pkg/net/http/#ListenAndServe), so long as `s` has to implement [http.Handler](https://golang.org/pkg/net/http/#Handler) and that allows us to bind that address to that server.

### The server file

The interesting part of that is `NewServer`, defined in `server.go`.
In that function we create a `server` struct with the database passed in (remember our main function passed in `NewMemoryStore()`), as well as a new router.
With the router we bind 2 [handler functions](https://golang.org/pkg/net/http/#HandleFunc) for our `/quotes` and `/random` functions and return the server.
Below the psuedo constructor, our server implements `ServeHTTP`, we will discuss how that works further on, the short answer being interfaces.

At the top of this file there is a declaration of a `Quote` struct, containing the text of the quote, the author, and an ID of the quote.
These fields also have a json [tag](https://medium.com/golangspec/tags-in-golang-3e5db0b8ef3e), these allow us to easily serialise and deserialise this data.

The `Datastore` declaration is an interface, for more reading on interfaces try [the tour of go](https://tour.golang.org/methods/9) or [go by example](https://gobyexample.com/interfaces).
Our interface declares three methods, `Insert`, `List`, and `Random`, this enables us compile time guarantees that anywhere we use `Datastore` we have access to those three methods.
Looking back at our `server` struct, it holds a `Datastore` field, function on the `server` struct can use any of these functions.

Last thing in this file is the handler functions themselves, `quote` and `randomQuote`.
First thing is these are both functions on `s`, this allows us to use our `Datastore` interface from earlier.
The `quotes` function accepts both `POST` and `GET` methods, so we need to switch based on the request info.
In the `POST` we [json decode](https://golang.org/pkg/encoding/json/) the users input, call our datastores Insert function - which we know must exist as the application has compiled.
After that we [marshal](https://golang.org/pkg/encoding/json/#Marshal) the data, and write that to our [`ResponseWriter`](https://golang.org/pkg/net/http/#ResponseWriter).
In the get, we call our datastores List function, again marshal it and write out the bytes to our `ResponseWriter`.
The `randomQuote` function behaves similarly; calling datastores Random and marshalling it.

## The memory file

The entry point here is the `NewMemoryStore()` function, it takes no arguments and returns a `memory` struct.
The `memory` struct contains a map of strings to `Quotes`, allowing us to store a lookup from IDs to Quotes.
When initialising the `quotes` field in memory, we use Go's built in [make](https://golang.org/pkg/builtin/#make), if we did not do this, we would end up assigning a key/value pair to a nil map resulting in a panic.

To comply with the `Datastore` interface described in the server file our `memory` struct must implement the `Insert`, `Random`, and `List` functions.
The `Insert` function uses the `random.go` file to create a random 8 character key ID and assigns that ID to the quote, this function makes use of the [rand](https://golang.org/pkg/math/rand/) package.
The `Random` function iterates the map, returning the first element, or an error if there are no elements in the map.
The `List` function iterates through the map returning it as a list.

## Exercises

- Try adding a `Get` function to the `Datastore` interface that allows a user to retrieve a quote by a specific ID
- Try implementing a datastore using local file storage using the [ioutil](https://golang.org/pkg/io/ioutil/), and [os](https://golang.org/pkg/os/)
- Try implementing a datastore using a database using the [sql](https://golang.org/pkg/database/sql/) packages (drivers need to be imported separately).

