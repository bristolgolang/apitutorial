# Lesson Web App

This lesson aims to show how to create a webpage entirely in go, with two examples; a classic todo page, and a quote of the day page.

There's exercises and further reading at the bottom of the page. To get the best out of this, we would suggest building your own simple webapp in parallel to this tutorial.

## How to run

To run in a terminal (this should be run from the root of this git repo)

```sh
go run lesson-webapp/main.go
```

Visit either `localhost:8080/quote` or `localhost:8080/todo` in your web browser

## Overview

How this application works in a nutshell is loading some html files with basic templating (similar to jinja 2, or other frameworks) into memory, having a router accept a path, create a data structure, render the template with the given data.

## [html/template](https://golang.org/pkg/html/template/)

A Go html template is a data driven html page. There's a lot of different formating actions available here, including for loops, conditionals. These are all based on, and better documented within the [text/template](https://golang.org/pkg/text/template/#hdr-Actions) package.

We have two examples in our lesson; a simpler html page for our quotes page where we only grab data from within our flat struct; and a todo list page which explores some of the actions such as range and conditionals.

Exploring the formatting directives on our quotes page we can see an example of `{{.Text}}`, the `.` here is a reference to the data passed in to the page, and `Text` is a reference to a field within the struct called `Text`. This means the data struct we pass in to has to be a flat struct with a field called `Text`.

The todopage is a little more interesting in that we range over `.Items`, which means our top level struct has to have an array called `Items`. Within the block of code inside of `range`, the value of `.` becomes an individual element within the `Items` slice.

## Creating the data

At the top of our `main.go` file we have structs that map to the directives within our template files. Looking out our `todoPage` func, we create a struct of our `TodoPageData` type, which we pass straight into our template using the [Execute](https://golang.org/pkg/text/template/#hdr-Actions) function.

A more interesting example is the `quote` function which creates an array of quotes, however our quote page only takes in a single quote, so we pick one at random using the [rand](https://golang.org/pkg/math/rand/) package.

## Exercises

- Create your own page
- Explore what happens when the struct passed in to `Execute` doesn't match the template.
- Combine this lesson with [making a downstream call to another API](../lesson-third-party/README.md)

## Further Reading

- [buffalo](https://github.com/gobuffalo/buffalo) - A Go web development eco-system, designed to make your project easier.
- [hugo](https://gohugo.io/) - The world’s fastest framework for building websites
- [gopheracadamy](https://blog.gopheracademy.com/advent-2017/using-go-templates/) - A lengthier tutorial
