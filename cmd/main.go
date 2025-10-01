package main

import (
	"log"
	"net/http"
	"os"

	"url_shortner/cmd/service_provider"
)

func main() {
	sp := service_provider.NewServiceProvider()

	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":8081"
	}

	log.Println("API server listening on", addr, "storage:", os.Getenv("STORAGE_MODE"))

	if err := http.ListenAndServe(addr, sp.GetRoutes()); err != nil {
		log.Fatal(err)
	}
}
