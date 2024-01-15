package main

import (
	"fmt"
	"net/http"

	handler "github.com/bladev/goshorty/Handler"
)


func main() {



	mux := defaultMux()
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