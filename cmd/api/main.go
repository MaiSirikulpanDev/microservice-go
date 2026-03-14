package main

import (
	"log"
	"microservice-go/internal/controller"
	"microservice-go/internal/service"
)

func main() {
	svc := service.NewCatFactService("https://catfact.ninja/fact")
	svc = service.NewLoggingService(svc)

	apiServer := controller.NewApiServer(svc)
	log.Fatal(apiServer.Start(":8080"))
}
