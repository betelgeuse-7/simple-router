package main

import (
	"fmt"
	"log"
	"net/http"
)

// Alias for the http.HandlerFunc type (why?,  idk.)
type myHandlerType func(w http.ResponseWriter, r *http.Request)

type Route struct {
	Path    string
	Handler myHandlerType
}

// the actual router type
type myRouter struct {
	Routes []Route
}

/**
* Return a new myRouter instance.
 */
func New() *myRouter {
	return &myRouter{}
}

/**
* Register a route with a callback function that handles the request
 */
func (mr *myRouter) RegisterRoute(path string, handler myHandlerType) {
	// pushing new Route to mr.Routes
	mr.Routes = append(mr.Routes, Route{Path: path, Handler: handler})
}

/**
* This method is required for myRouter type to implement the http.Handler interface.
*
* To use the router as a "handler", we need to pass myRouter as a second argument
* to http.ListenAndServe() function.
 */
func (mr myRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	routes := mr.Routes

	var handlerToRun myHandlerType

	for _, ro := range routes {
		if path == ro.Path {
			handlerToRun = ro.Handler
		}
	}

	// handlerToRun is nil, if we couldn't match
	// the 'path' with any route in mr.Routes
	if handlerToRun == nil {
		// in that case, we return '404 Not Found'
		fmt.Fprintf(w, "404 Not Found")
	} else {
		handlerToRun(w, r)
	}
}

func main() {
	// initialize the router
	var router *myRouter = New()

	router.RegisterRoute("/cat", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "CATS")
	})

	router.RegisterRoute("/dog", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "DOGS")
	})

	fmt.Println("Routes: ", router.Routes)

	log.Fatalln(http.ListenAndServe(":8000", router))
}
