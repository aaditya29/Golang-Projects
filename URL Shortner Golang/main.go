package main

type Config struct { // Config is a struct describing the configuration which are parsed from command line arguments
	PathToYAML string
	PathToJSON string
}

func main() {
	config := getConfig()

	yamlBytes := getFileBytes(config.PathToYAML)
	jsonBytes := getFileBytes(config.PathToJSON)

	mux := makeDefaultMux()
	mapHandler := makeMapHandler(mux)

	handler := mapHandler
	if yamlBytes != nil {
		handler = makeYAMLHandler(yamlBytes, &mapHandler)
	} else if jsonBytes != nil {
		handler = makeJSONHandler(jsonBytes, &mapHandler)
	}
	startServer(handler)
}
