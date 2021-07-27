// Adding virs REST API Handlers
type Virus struct {
	Types       string `json:"Types"`
	Description string `json:"description"`
}

var virus []Virus

//Defining handler to fetch the details of all the viruses
func getVirusHandler(w http.ResponseWriter, r *http.Request) {
	//Converting the virus variable to json
	virusListBytes, err := json.Marshal(virus)

	if err != nil { //if error occurs
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Writing list of viruses to the response
	w.Write(virusListBytes)
}

// Handler to create new entry of viruses
func createVirusHandler(w http.ResponseWriter, r *http.Request) {
	//Creating new instance of virus
	virus := Virus{}

	//Sending all our data as HTML form data
	err := r.ParseForm() //the `ParseForm` method of the request, parses the form values

	if err != nil { //if error occurs
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Get the information about the virus from the form info
	virus.types = r.Form.Get("types")
	virus.Description = r.Form.Get("description")

	// Append our existing list of virus with a new entry
	viruses = append(viruses, virus)

}