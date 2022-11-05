package main

import (
	"context"
	"fmt"
	"log"
	"logservice/registry"
	"net/http"
)

func main() {
	http.Handle("/service",
		&registry.RegistryService{})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var srv http.Server
	srv.Addr = registry.ServerPort

	go func() {
		log.Println(srv.ListenAndServe())
		cancel()
	}()
	go func() {
		fmt.Printf("Registry service started. Press any key to stop")
		var s string
		fmt.Scanln(&s)
		srv.Shutdown(ctx)
		cancel()
	}()
	<-ctx.Done()
	fmt.Println("Shutting down registryservice")
}
