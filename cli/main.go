package main

import (
	"log"
	"net/http"

	"github.com/Jarigyani/go_ogp"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env.yaml.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	http.HandleFunc("/ogp", go_ogp.GetOGP)
	http.ListenAndServe(":8080", nil)
}
