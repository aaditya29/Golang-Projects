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
}