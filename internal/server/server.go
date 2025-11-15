package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Xebec19/simple-bank/internal/controllers"
	"github.com/Xebec19/simple-bank/internal/db"
	"github.com/Xebec19/simple-bank/internal/routes"
	"github.com/gin-gonic/gin"
)

type Server interface {
	Start(address string) error
}

type server struct {
	router *gin.Engine
	query  *db.Queries
}

func NewServer(ctx context.Context) (Server, error) {
	// Initialize database connection
	query, err := db.GetDB(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	router := gin.Default()

	s := &server{
		router: router,
		query:  query,
	}

	s.registerRoutes()

	return s, nil
}

func (s *server) registerRoutes() {
	// Health check endpoint
	s.router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// API v1 routes
	v1 := s.router.Group("/api/v1")

	// Initialize controllers
	accountController := controllers.NewAccountController(s.query)

	// Register route groups
	routes.RegisterAccountRoutes(v1, accountController)
}

func (s *server) Start(address string) error {
	return s.router.Run(address)
}
