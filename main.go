package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	r := mux.NewRouter()
	port := getPort()

	r.HandleFunc("/", indexHandler).Methods("GET")

	r.HandleFunc("/webhook", webhookGetHandler).Methods("GET")

	fmt.Printf("Server up and running on port: %s\n", port)

	err := http.ListenAndServe(port, r)

	if err != nil {
		log.Fatal("Error listening and server:", err)
	}
}

func indexHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Got my server up and running in Go.")
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3500"
		fmt.Printf("Port not defined. Using default port %s\n", port)
	}
	return port
}
