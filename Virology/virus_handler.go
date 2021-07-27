// Adding virs REST API Handlers
type Virus struct {
	Types       string `json:"Types"`
	Description string `json:"description"`
}

var virus []Virus