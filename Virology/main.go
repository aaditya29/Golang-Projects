// code to start the server
package main

import "net/http"

func main() {
	http.HandleFunc("/", handler) // The "HandleFunc" method accepts a path and a function as arguments

	http.ListenAndServe(":8080", nil)
}
