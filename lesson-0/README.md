# Lesson 0

This lesson aims to demonstrate an API at its simplest.

## How to run

To run in a terminal

```sh
go run main.go
```

Alternatively there you can build run it as a docker image

```sh
docker build . -t apitutorial
docker run -p 8080:8080 apitutorial:latest
```

## How it works

The [net/http](https://golang.org/pkg/net/http) package allows us to create a [handler function](https://golang.org/pkg/net/http/#Handle) which takes a pattern, as well as a handler function.
If the request to the server matches the pattern then the handler function will be called.

Our handler function takes both a [http.ResponseWriter](https://golang.org/pkg/net/http/#ResponseWriter), which we can write our response to, as well as a [*http.Request](https://golang.org/pkg/net/http/#Request) where we can gather more information from the original request into the server, such as path parameters, query parameters, etc.

To send a response all we need to do is Write a series of bytes to our ResponseWriter. We use the code `w.Write([]byte("Hello, world!"))`.

## Docker

Install [docker](https://docs.docker.com/install/) for this bit.

We can dockerise the service too, more than that it is possible to have a docker image contain only the compiled Go binary, with nothing else (either using scratch with nothing else, or distroless with just enough to have a healthy web service).
To do this, we build our service with a Go image, then we can copy the binary over to a distroless docker image.

We can run this container in kubernetes, as a serverless container, or however else we see fit. Everything it needs is within the container, making it super portable.

Try typing `docker images` on the command line, search for our image and check the size.
Compare that to the size of the Go docker image used to build the service.

## Exercises

- See if you can parse a query or path parameter, and respond with a name. For example, `curl localhost:8080/developer` may respond with `Hello, developer!`.
