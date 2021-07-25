//testing handler
package main

import (
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) { // In this function, forming a new HTTP request which is going to be passed to our handler.

	req, err := http.NewRequest("GET", "", nil) // here first argument is the method, the second argument is the route (which we are leaving blank for a while)
	//and the third is the request body, which we don't have in this case.
}
