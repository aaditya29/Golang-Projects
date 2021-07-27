//testing handler
package main

import (
	"io/ioutil"
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
	mockServer := httptest.NewServer(r) // creating a new server for testing purpose
	//The mockServer will run a server and exposes its location in the URL attribute
	resp, err := http.Get(mockServer.URL + "/hello") // making a GET request to the "hello" route we defined in the router

	if err != nil { // For handling error
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK { // Code should be '200' if status is OK.
		t.Errorf("Status should be ok, got %d", resp.StatusCode)
	}

	//Reading the response body and converting into string now
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body) //Reading the body into a collection of bytes (b)
	if err != nil {
		t.Fatal(err)
	}

	// convert the bytes to a string
	respString := string(b)
	expected := "Hello World!"

	if respString != expected {
		t.Errorf("Response should be %s, got %s", expected, respString) // Checking if response is expected i.e. Hello World. If it comes as expected then route is correct
	}
}

func TestRouterForNonExistentRoute(t *testing.T) { //function for testing if we hit any other route than GET/hello
	r := newRouter()
	mockServer := httptest.NewServer(r)
	//Making request to route we didn't define before this function for example POST /hello route
	resp, err := http.Post(mockServer.URL+"/hello", "", nil)

	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Status should be 405, got %d", resp.StatusCode) //Method 405 is for if method is not allowed
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	respString := string(b)
	expected := "" //Expecting empty sting this time unlike in previous function where expected response was Hello World

}

//Testing static file server
func TestStaticFileServer(t *testing.T) {
	r := newRouter()
	mockServer := httptest.NewServer(r)

	//We want to go in `GET /assets/` route to get the index.html file response
	resp, err := http.Get(mockServer.URL + "/assets/")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be 200, got %d", resp.StatusCode) //Keeping status '200' i.e. OK
	}

	contentType := resp.Header.Get("Content-Type")    //HTML File can be huge to test so we do testing on the
	expectedContentType := "text/html; charset=utf-8" //content-type header is "text/html; charset=utf-8" making sure that an html file has been served

}
