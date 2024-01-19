package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"gopkg.in/yaml.v3"
)

//This returns 'hello world' to on 'localhost:port/'
func DefaultHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w , "Hello, Word")
}


// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	//	TODO: Implement this...
	return func(w http.ResponseWriter, r *http.Request) {
		path, ok := pathsToUrls[r.URL.Path]

		if ok {
			http.Redirect(w, r, path, http.StatusPermanentRedirect)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}


type URLMapper struct {
	Path string `yaml:"path" json:"path"` 
	Url string `yaml:"url" json:"url"`
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// TODO: Implement this...
	var mapper []URLMapper

	err := yaml.Unmarshal(yml, &mapper)

	if err != nil {
		return nil, err
	}
	
	return func(w http.ResponseWriter, r *http.Request){

		url := strings.Split(r.URL.Path, "/")
		path := "/" + url[len(url) - 1]

		for _, mapp := range mapper {
			if mapp.Path == path {
				http.Redirect(w, r, mapp.Url, http.StatusMovedPermanently)
				return
			}
		}
		fallback.ServeHTTP(w, r)

	}, nil
}


func JsonHandler(jsn []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var mapper []URLMapper

	err:= json.Unmarshal(jsn, &mapper)

	if err != nil {
		return nil, err
	}

	
	return func(w http.ResponseWriter, r *http.Request){
	url := strings.Split(r.URL.Path, "/")
	path := "/" + url[len(url) - 1]
	for _, mapp := range mapper {
		if mapp.Path == path {
			http.Redirect(w, r, mapp.Url, http.StatusMovedPermanently)
			return
		}
	}
	fallback.ServeHTTP(w, r)
	}, nil
}