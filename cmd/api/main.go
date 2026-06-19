package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	apserver "github.com/maneul0498-netizen/unicomer_tech_challenge/internal/server"
)

// @title Unicomer tech challenge
// @version 1.0
// @description Tech Challenge
// @host localhost:8080
// @BasePath /api/v1/
func main() {
	// Canal de señales del SO
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(quit)

	// Crear servidor
	s, err := apserver.New()
	if err != nil {
		log.Fatalf("unable to create server: %v", err)
	}

	// Canal de errores del server
	errCh := make(chan error, 1)

	// Levantar servidores (no bloquea)
	s.Run(errCh)

	log.Println("application started")

	// Esperar señal o error del server
	select {
	case sig := <-quit:
		log.Printf("shutdown signal received: %s", sig.String())

	case err := <-errCh:
		log.Printf("server error: %v", err)
	}

	// Contexto de apagado
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("shutting down servers...")

	if err := s.Shutdown(ctx); err != nil {
		log.Printf("graceful shutdown failed: %v", err)
	} else {
		log.Println("server stopped gracefully")
	}
}
