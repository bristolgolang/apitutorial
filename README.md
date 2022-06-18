# API Tutorial

Hello! This workshop will focus on creating APIs using just the core library of Go.

## How to learn

To get the most out of this guide, here are a couple of suggestions;
- Follow the READMEs, exploring the code and playing with the projects you see here
- Create your own project and copy this guide to recreate the lessons
- Clone/fork the repo and try tweaking/breaking things
- Follow the READMEs, but do your own thing, use different APIs, datastores, etc.

## Lessons

These lessons are largely approachable in any order, provided the first lesson is understood.

### [Lesson 0](./lesson-0/README.md)

This serves as a pre-cursor to the rest of the lessons, it sets up a basic hello world server.

### [Lesson - Using third party APIs](./lesson-third-party/README.md)

This lesson aims to teach how to write APIs that use other APIs. Useful for writing slackbots, logic APIs that call data layer APIs, or anything else that catches your interest.

### [Lesson - Creating Webapps](./lesson-webapp/README.md)

This lesson aims to teach how to write webapps, to render HTML pages using a Go server.

### [Lesson - Using datastores](./lesson-datastore/README.md)

This lesson aims to teach how to write APIs that need to work with storage; in memory, file storage, databases, or anything else. It's recommended to understand [serialising/deserialising](https://golang.org/pkg/encoding/json/), which is discussed in the third party API lessons.

### [Lesson - httptest](./lesson-httptest/README.md)

This lesson explores the ways in which we can test our http server.

### [Lesson - middleware](./lesson-middleware/README.md)

This lesson explores the ways in which we can add middleware to our http server.

## Further reading

- [Writing Web Applications](https://golang.org/doc/articles/wiki/) on the golang docs.
- [Go by example](https://gobyexample.com/http-servers).
- [OpenAPI code gen tools](https://github.com/deepmap/oapi-codegen) if you're supporting openAPI/swagger.
- [Dockerize](https://blog.golang.org/docker) your Go binaries.

