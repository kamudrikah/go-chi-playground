package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mudrikah/go-chi-playground/server"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	// Start HTTP server
	go func() {
		defer wg.Done()
		r := chi.NewRouter()
		r.Use(middleware.Logger)
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Welcome - HTTP & gRPC Server"))
		})

		log.Println("Starting HTTP server on :3000")
		if err := http.ListenAndServe(":3000", r); err != nil {
			log.Fatalf("Failed to start HTTP server: %v", err)
		}
	}()

	// Start gRPC server
	go func() {
		defer wg.Done()
		if err := server.StartGRPCServer(50051); err != nil {
			log.Fatalf("Failed to start gRPC server: %v", err)
		}
	}()

	wg.Wait()
}
