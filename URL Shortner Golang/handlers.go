package main

import "net/http"

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	/*
	   Here, function MapHandler will return an http.HandlerFunc (which also implements another function http.Handler)
	   that will attempt to map any paths (keys in the map) to their corresponding URL (values) that each key in the map points to, in string format).
	*/

	// If the path is not provided in the map, then the fallback http.Handler will be called instead.
	return func(w http.ResponseWriter, r *http.Request) {
		originalURL, ok := pathsToUrls[r.URL.Path]
		if ok {
			http.Redirect(w, r, originalURL, 301)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}

}

func YAMLHandler(YAML []byte, fallback http.Handler) (http.HandlerFunc, error) {
	/*
		Here, YAMLHandler will parse the provided YAML and then return an http.HandlerFunc (which also implements http.Handler)
		that will attempt to map any paths to their correspondingURL.
		If the path is not provided in the YAML, then the fallback http.Handler will be called instead.
	*/

	/*
		YAML is expected to be in the format:
		path: /some-path
		url: https://www.some-url.com/demo
	*/

}
