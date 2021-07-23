package main

import "flag"

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

func getConfig() *Config {
	config := Config{}
	flag.StringVar(&config.PathToYAML, "yaml", "", "--yaml=path/to/file.yml")
	flag.StringVar(&config.PathToJSON, "json", "", "--json=path/to/file.json")
	flag.Parse()
	return &config
}
