package main

import (
	"net/http"

	"github.com/aluyapeter/salon-finder-backend/internal/config"
	"github.com/aluyapeter/salon-finder-backend/internal/handlers"
	"github.com/aluyapeter/salon-finder-backend/internal/logger"
)

// func healthHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Salon Finder API v0.1")
// }

func main() {
	cfg := config.Load()

	log := logger.New()
	// http.HandleFunc("/health", healthHandler)
	mux := http.NewServeMux()
	mux.HandleFunc("/health", handlers.Health)

	log.Info("Starting server on port " + cfg.Port)

	// fmt.Println("Starting server on :8080...")

	err := http.ListenAndServe(":"+cfg.Port, mux)
	if err != nil {
		log.Error("Error starting server: " + err.Error())
	}
}