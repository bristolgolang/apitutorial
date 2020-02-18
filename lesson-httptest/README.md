# Lesson HTTPTest

Aims to teach how to test an API at the server layer, even through the routing layer through [httptest](https://golang.org/pkg/net/http/httptest/)

## How it works

Firstly, looking at the `main.go` file, we've split out a `newServer` function to create ourselves a server function that handles routing. To understand this concept further, either explore the lesson on datastores, or read up on [http.ListenAndServer](https://golang.org/pkg/net/http/#ListenAndServe) in particular the [http.Handler](https://golang.org/pkg/net/http/#Handler) argument.

This exposes a couple of things we can test, we can either expose our `helloWorld` function directly, or we can test the server and routing are set up correctly. Which level of testing to do in your own project is your choice.

### Directly testing the Handler

The goal here is to test the functionality of the handler function itself, and not to worry if it is being routed correctly at all.
To test our function, we need to send a request, so we need to call the [NewRequest](https://golang.org/pkg/net/http/#NewRequest) function in the net/http package.
If this errors, our test is not going to go well, so we may as well call [t.Fatal](https://golang.org/pkg/testing/#B.Fatal), this stops our test immediately and fails it.

Provided the request was made successfully, we need a way of recording what was written to the handlers `ResponseWriter`, we use httptests [NewRecorder](https://golang.org/pkg/net/http/httptest/#NewRecorder) for this.
This struct will hold data about what would be responded, like the status code and the response body.

After this, we can create a handler on our function we want to test using [HandlerFunc](https://golang.org/pkg/net/http/#HandlerFunc).
We can then serve the request, with the response recorder to the handler we just made.
This is where our `helloWorld` function is executed.

After this, we can do regular assertions like we would with any other unit test.
We can assert the status code matched what we wanted, or the Body has the text we expect.

### Testing the routing

This test ensures not only our function works correctly, but that the routing on the server is wired up correctly.
By its nature, this test is potentially more brittle, but exposes another layer of potential bugs.

To execute this test, we use httptests [NewServer](https://golang.org/pkg/net/http/httptest/#NewServer) function, using the response from `newServer()` (remember that that was the function where we did the wiring up of the routing).
We also need to be careful about closing this new test server down by deferring `s.Close()`

We create a request in a similar fashion to directly testing the handler func above, by calling http.NewRequest.
The slight caveat here is, we must set the URL to be equal to the URL of our test server + whatever path we want to test, so our path looks like `s.URL + "/"`.
We then use the [default client](https://golang.org/pkg/net/http/#Client) to [Do](https://golang.org/pkg/net/http/#Client.Do) the request we want to test.

Finally we can call whatever assertions we want on the response, checking the status code, and the body etc.

## Exercises

- With each test, try changing the routing on the request, what do you expect to happen and why?
- Expand the routing test to explore [NotFound](https://golang.org/pkg/net/http/#pkg-constants) or other responses.
- Try a more complex test that manipulates the input data, and assert the value you get out is correct (try saying hello to someone using a query parameter)
