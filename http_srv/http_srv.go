package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	responseString := `{
		"id": "1",
		"message": "Hello world"
	}`
	fmt.Fprintf(w, responseString)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}