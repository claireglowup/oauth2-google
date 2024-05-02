package main

import (
	"google-oauth2/google"
	"log"
	"net/http"
)

func main() {

	googleOauth := google.GoogleConfig()

	http.HandleFunc("/auth", googleOauth.GoogleLogin)
	http.HandleFunc("/auth/callback", googleOauth.GoogleCallback)

	port := ":8080"
	log.Printf("Server starting at localhost:%s", port)
	log.Fatal(http.ListenAndServe(port, nil))

}
