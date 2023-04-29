package main

import "urlshortener/shortener"

func main() {
	// mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := shortener.MapHandler(pathsToUrls, nil)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yaml := `
	- path: /urlshort
	  url: https://github.com/gophercises/urlshort
	- path: /urlshort-final
	  url: https://github.com/gophercises/urlshort/tree/solution
	`

	yamlHandler, err := urlshort.YAMLHandler

}

func defaultMux() {
	panic("unimplemented")
}
