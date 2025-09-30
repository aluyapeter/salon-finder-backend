package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aluyapeter/salon-finder-backend/internal/config"
	"github.com/aluyapeter/salon-finder-backend/internal/db"
	"github.com/aluyapeter/salon-finder-backend/internal/handlers"
	"github.com/aluyapeter/salon-finder-backend/internal/logger"
)

// func healthHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Salon Finder API v0.1")
// }

func main() {
	cfg := config.Load()
	logger := logger.New()

	database, err := db.Connect(cfg.DBUrl)
	if err != nil {
		logger.Error("failed to connect to DB: " + err.Error())
		return
	}
	defer database.Close()

	// http.HandleFunc("/health", healthHandler)
	mux := http.NewServeMux()
	mux.HandleFunc("/health", handlers.Health)

	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: mux,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		logger.Info("Starting server on port " + cfg.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("server error: " + err.Error())
		}
	}()

	<-stop
	logger.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Error("Forced to shutdown: " + err.Error())
	} else {
		logger.Info("server stopped gracefully")
	}

	// fmt.Println("Starting server on :8080...")

	// err := http.ListenAndServe(":"+cfg.Port, mux)
	// if err != nil {
	// 	logger.Error("Error starting server: " + err.Error())
	// }
}
