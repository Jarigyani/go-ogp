package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ogp", getOGP)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
