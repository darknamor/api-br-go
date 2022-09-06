package main

import (
	"api-br-go/routes"
	"log"
	"net/http"
)

func main() {
	router := routes.NewRouter()

	server := http.ListenAndServe(":8080", router)
	log.Fatal(server)
}
