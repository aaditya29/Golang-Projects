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

	//Declaring static file directory and pointing to it
	staticFileDirectory := http.Dir("./assets/")
	//Declaring handler that router requsts with filename
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory)) // The fileserver is wrapped in the `stripPrefix` method
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")                      //The "PathPrefix" method acts as a matcher, and matches all routes starting with "/assets/", instead of the absolute route itself

	//Adding virusHandlers to our router
	r.HandleFunc("/virus", getVirusHandler).Methods("GET")
	r.HandleFunc("/virus", createVirusHandler).Methods("POST")
	return r
}

func main() {

	// Old router is deleted and new router is now formed by calling the `newRouter` constructor function that we defined above this function.
	r := newRouter()
	http.ListenAndServe(":8080", r)
}

func handler(w http.ResponseWriter, r *http.Request) { // handler is the handler function. It has to follow the function signature of a ResponseWriter and Request type as the arguments.
	fmt.Fprintf(w, "Hello World!") //Writing Hello World in the response writer.
}
