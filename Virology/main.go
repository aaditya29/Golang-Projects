// code to start the server
package main

import "net/http"

func main() {
	http.HandleFunc("/", handler) // The "HandleFunc" method accepts a path and a function as arguments

	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) { // handler is the handler function. It has to follow the function signature of a ResponseWriter and Request type as the arguments.

}
