package main

import (
	"log"
	"net/http"
	"os"

	"github.com/building-microservices-with-go/chapter4/data"
	"github.com/building-microservices-with-go/chapter4/handlers"
)

func main() {
	serverURI := "localhost"
	if os.Getenv("DOCKER_IP") != "" {
		serverURI = os.Getenv("DOCKER_IP")
	}

	store, err := data.NewMongoStore(serverURI)
	if err != nil {
		log.Fatal(err)
	}

	handler := handlers.Search{DataStore: store}
	err = http.ListenAndServe(":8323", &handler)
	if err != nil {
		log.Fatal(err)
	}
}
