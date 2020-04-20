// This is a single line comment
/*
This comment
is over
multiple lines
*/
package main

// listing any packages that we need to import
import (
	// This time we are using some more packages
	"fmt"
	"log"
	"net/http"
)

// The main function of the package does not take any values in or return any
func main() {
	// A lot of this is adapted to what we need from the official golang documentation https://golang.org/doc/articles/wiki/

	// This will print "I am alive" to stdout
	fmt.Println("I am alive")
	/*  We have to add capalities to handle http requests
	    therefore we will make use of the net/http package.
	    This allows us to make a simple webserver

	*/
	// Call to http.HandleFunc, telling it to handle all requests to "/" using the "handler" function
	http.HandleFunc("/", handler)
	// Log a fatal in case something goes wrong with the "Listen and Serve"
	// Listen and serve is specified to listen on port 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// The function handler is of the type http.HandlerFunc. It takes an http.ResponseWriter and an http.Request as its arguments.
// An http.ResponseWriter value assembles the HTTP server's response; by writing to it, we send data to the HTTP client
func handler(w http.ResponseWriter, r *http.Request) {
	/*  An http.Request is a data structure that represents the client HTTP request. r.URL.Path is the path component of the request URL.
	The trailing [1:] means "create a sub-slice of Path from the 1st character to the end."
	This drops the leading "/" from the path name.*/
	fmt.Fprintf(w, "Hi there, I am alive at /%s!", r.URL.Path[1:])
}
