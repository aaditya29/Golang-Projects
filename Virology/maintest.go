//testing handler
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) { // In this function, forming a new HTTP request which is going to be passed to our handler.

	req, err := http.NewRequest("GET", "", nil) // here first argument is the method, the second argument is the route (which we are leaving blank for a while)
	//and the third is the request body, which we don't have in this case.

	if err != nil { //if error occurs
		t.Fatal(err)
	}

	/*
		Usinghttptest library to create an http recorder.
		This recorder will act as the target of our http request.
		In layman terms it works as a mini-browser, which will accept the result ofthe http request that we make)
	*/
	recorder := httptest.NewRecorder()

	hf := http.HandlerFunc(handler) //creating an HTTP handler from our handler function from file which we want to test.

	hf.ServeHTTP(recorder, req) // serving the HTTP Request to recorder. This line tests the handler we wanna test

	if status := recorder.Code; status != http.StatusOK { // checking STATUS of expected code
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Checking if response is expected one or not
	expected := `Hello World!`
	actual := recorder.Body.String()
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

//Testing our routing
func TestRouter(t *testing.T) { // Instantiating the router using the constructor function that we defined above
	r := newRouter()

}
