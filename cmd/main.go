package main

import (
	"fmt"
	"log"
	"net"
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

	grpcAddr := os.Getenv("GRPS_ADDR")
	if grpcAddr == "" {
		grpcAddr = ":50051"
	}

	go func() {
		log.Println("API server listening on", addr, "storage:", os.Getenv("STORAGE_MODE"))
		if err := http.ListenAndServe(addr, sp.GetHTTPServer()); err != nil {
			log.Fatal(err)
		}
	}()
	grpcServer := sp.GetGRPCServer()

	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	if config := grpcServer.Serve(listen); config != nil {
		fmt.Println("failed to serve: %v\n", listen)
	}
}
