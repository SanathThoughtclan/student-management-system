package app

import (
	"context"
	"log"
	"net/http"
	"student-management-system/config"
	"student-management-system/routes"
	"student-management-system/utils"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Server struct {
	server   *http.Server
	database *mongo.Database
}

func NewServer() *Server {
	cfg := config.LoadConfig()

	clientOptions := options.Client().ApplyURI(cfg.Database.URI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	db := client.Database(cfg.Database.Database)

	router := routes.NewRouter(db)

	return &Server{
		server: &http.Server{
			Addr:    ":" + cfg.Server.Port,
			Handler: router,
		},
		database: db,
	}
}

func (s *Server) Run() error {
	utils.InitLogger()
	log.Println("Starting server on", s.server.Addr)
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown() error {
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		return err
	}

	log.Println("Server shut down gracefully")
	return nil
}
