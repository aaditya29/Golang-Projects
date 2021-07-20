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
