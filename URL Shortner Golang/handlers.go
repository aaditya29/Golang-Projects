package main

import "net/http"

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	/*
	   Here, function MapHandler will return an http.HandlerFunc (which also implements another function http.Handler)
	   that will attempt to map any paths (keys in the map) to their corresponding URL (values) that each key in the map points to, in string format).
	*/

}
