package main

import (
	"flag"
	"fmt"
	"net/http"

	sqlDB "github.com/bladev/goshorty/Database"
	handler "github.com/bladev/goshorty/Handler"
)

var jsonFile, yamleFile, dbName string

func main() {

	flag.StringVar(&jsonFile, "json", "", "takes the path to json file")
	flag.StringVar(&yamleFile, "yaml", "", "takes the path to json file")
	flag.StringVar(&dbName, "db", "", "takes a database name and creates a db if one is not provided")

	flag.Parse()
	mux := defaultMux()

	h := defaultMapper(mux)
	if yamleFile != "" {

		mux.HandleFunc("/yaml/", yamlHandler(yamleFile, h))
		fmt.Println("yaml file added!")
	} 
	
	if jsonFile != "" {

		mux.HandleFunc("/json/",jsonHandler(jsonFile, h))
		fmt.Println("json file added!")
	} 
	
	if dbName != "" {
		db := &sqlDB.Database{
			DBName: dbName,
		}

		db.Init()
		
		mux.HandleFunc("/db/", dbHandler(db, h))
		fmt.Println("database added!")
		
	}

	startServer(h)
}

func startServer(handler http.HandlerFunc) {
	fmt.Println("Starting server...")
	http.ListenAndServe(":8080", handler)
	
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.DefaultHandler)

	return mux
}

// Handlers

func defaultMapper(mux http.Handler) http.HandlerFunc {
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	return handler.MapHandler(pathsToUrls, mux)

}

func yamlHandler(path string, mux http.Handler) http.HandlerFunc {
	data, err := handler.YAMLHandler(handler.ReadFile(path), mux)

	if err != nil {
		panic(err)
	}
	return data
}

func jsonHandler(path string, mux http.Handler) http.HandlerFunc {
	data, err := handler.JsonHandler(handler.ReadFile(path), mux)
	if err != nil {
		panic(err)
	}
	return data
}

func dbHandler(db *sqlDB.Database, mux http.Handler )http.HandlerFunc{
	return sqlDB.DBHandlerfunc(db, mux)
}