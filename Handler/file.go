package handler

import (
	"log"
	"os"
)


func ReadFile(name string) (data []byte) {

	data, err := os.ReadFile(name)

	if err != nil {
		log.Fatal(err)
	}

	return 
}