# Calling third party APIs

## How to run

To run from the root of this repository;

```sh
go run lesson-third-party/main.go
```

then from a different terminal session (or program like postman)

```sh
curl localhost:8080/quotes/search/oneself
curl localhost:8080/quotes/random
```

## Do it yourself

A few alternative APIs to challenge yourself;
- https://docs.spacexdata.com
- https://github.com/shevabam/breaking-bad-quotes
- https://github.com/skolakoda/programming-quotes-api
- https://openlibrary.org/dev/docs/api/books
- https://xkcd.com/json.html

Also, a [meta list](https://github.com/public-apis/public-apis
) of APIs if non of those take your fancy.

Try implementing your own functionality, for example;
- implement a 'random' functionality for an API that has a `/list` resource but no natural `/random`
- combine multiple requests, i.e. take a quote and use a sentiment analysis on it
- return a json object instead of plain text

## How it works

This section aims to walkthrough what is happening in this API.

### Handling input

In this particular example we only handle path parameters in our `getSearchQuote`, we can just as easily add query parameters, as well as handling URL forms with just the core library.

### Making requests

In our main function, we create a http client, with a default timeout of 10 seconds. This is in case our downstream server is unresponsive, our server will at least return to the caller.

Within both of our handler functions, we make a `GET` request which returns both a `*http.Response` and an `error`. Perform an error check, if there is an error, return a `500` with a user friendly error message. If there's no error, we can continue. It's important to note the body of the response is a `io.ReadCloser`, so we should defer a `Close()` on the body once we're sure it's usable.

We can use `io.Copy` to copy the contents of our `GET` to our `ResponseWriter` if we want to return the response we get verbatim.

### Serialization/Deserialization

The response we got from our `GET` request is a slice of bytes, to make this something useful within our own application it's important to serialize it into a struct we can use and manipulate. To do this, we create an instance of the response we're expecting `Quote` for a singular response, `Quotes` for our search query. We can then create a new decoder using the response body as it implements `io.Reader`. 

For quickly creating the structs, mholt has a fantastic resource called [json=to-go](https://mholt.github.io/json-to-go/). Chuck in a json response example, and use the generated structs, this service can fall down on traditional map types though.

### Responding to the caller

If all goes well, we can just use `w.Write` to write a slice of bytes out, in the example here we're created a `String` function that prints out the quote in a nice plain text format.

