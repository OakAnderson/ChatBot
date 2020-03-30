package main

import (
	"log"
	"net/http"

	host "github.com/OakAnderson/ChatBot/http"
)

func main() {
	log.Printf("Server running on %s\n", host.PORT)
	http.ListenAndServe(host.PORT, http.HandlerFunc(host.Handler))
}
