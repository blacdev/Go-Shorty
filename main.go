package main

import (
	"flag"
	"fmt"
	"net/http"

	handler "github.com/bladev/goshorty/Handler"
)


var jsonFile, yamleFile string
func main() {


	flag.StringVar(&jsonFile, "json", "", "takes the path to json file")
	flag.StringVar(&yamleFile, "yaml", "", "takes the path to json file")

	flag.Parse()
	mux := defaultMux()
	
	if yamleFile != ""{
		startServer(yamlHandler(yamleFile, mux))

	}

	if jsonFile != ""{
		startServer((jsonHandler(jsonFile, mux)))
	}
	h := defaultMapper(mux)
	
	startServer(h)
}


func startServer(handler http.HandlerFunc){
	fmt.Println("Starting server...")
	http.ListenAndServe(":8080", handler)
	
}


func defaultMux() *http.ServeMux{
	mux := http.NewServeMux()

	mux.HandleFunc ("/", handler.DefaultHandler)
	
	return mux
}


// Handlers

func defaultMapper(mux http.Handler) http.HandlerFunc{
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc": "https://godoc.org/gopkg.in/yaml.v2",
	}
	return handler.MapHandler(pathsToUrls, mux)

}

func yamlHandler(path string, mux http.Handler) http.HandlerFunc {
	data, err:= handler.YAMLHandler(handler.ReadFile(path), mux)

	if err != nil {
		panic(err)
	}
	return data
}

func jsonHandler(path string, mux http.Handler) http.HandlerFunc{
	data, err := handler.JsonHandler(handler.ReadFile(path), mux)
	if err != nil {
		panic(err)
	}
	return data
}