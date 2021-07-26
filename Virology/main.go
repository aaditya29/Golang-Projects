// code to start the server
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//Testing the router
func newRouter() *mux.Router { // Here we are creating new router function which creates the router and returns it to us.
	//This function will be used for instantiating and testing the router outside of the main function
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")
	return r
}

func main() {

	r := mux.NewRouter() //Declaring a new router

	r.HandleFunc("/hello", handler).Methods("GET") //This allows us to declare METHODS this path will be valid for

	http.HandleFunc("/", handler) // The "HandleFunc" method accepts a path and a function as arguments

	http.ListenAndServe(":8080", r)
}

func handler(w http.ResponseWriter, r *http.Request) { // handler is the handler function. It has to follow the function signature of a ResponseWriter and Request type as the arguments.
	fmt.Fprintf(w, "Hello World!") //Writing Hello World in the response writer.
}
