package main

type Config struct { // Config is a struct describing the configuration which are parsed from command line arguments
	PathToYAML string
	PathToJSON string
}

func main() {
	config := getConfig()

	yamlBytes := getFileBytes(config.PathToYAML)
	jsonBytes := getFileBytes(config.PathToJSON)
}
