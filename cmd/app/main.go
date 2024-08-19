package main

import (
	"log"
	"os"
	"os/signal"
	"student-management-system/internal/app"
	"syscall"
)

func main() {
	server := app.NewServer()

	go func() {
		if err := server.Run(); err != nil {
			log.Fatalf("Server execution interrupted : %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	if err := server.Shutdown(); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
}
