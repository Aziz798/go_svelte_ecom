package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"go_ecom/internal/database"
	"go_ecom/internal/routes/product"
	"go_ecom/internal/routes/user"

	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port       int
	db         database.Service
	httpServer *http.Server
}

func NewServer() *Server {
	// Retrieve port from environment variable
	portStr := os.Getenv("PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("Invalid port number: %v", err)
	}

	// Initialize the database service
	dbService := database.New()

	// Create the server instance
	server := &Server{
		port: port,
		db:   dbService,
	}

	// Initialize the HTTP server configuration
	server.httpServer = &http.Server{
		Addr:         fmt.Sprintf(":%d", server.port),
		Handler:      server.RegisterRoutes(), // Initialize routes
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}

// RegisterRoutes sets up the routes for the server.
func (s *Server) RegisterRoutes() http.Handler {
	r := mux.NewRouter()
	user.RegisterUserRoutes(r, s.db.DB())       // Pass the *sql.DB instance to the user routes
	product.RegisterProductRoutes(r, s.db.DB()) // Pass the *sql.DB instance to the product routes
	return r
}

// ListenAndServe starts the HTTP server and listens for requests.
func (s *Server) ListenAndServe() error {
	return s.httpServer.ListenAndServe()
}

// healthHandler is an example handler that returns the health of the database.
func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	health := s.db.Health()
	for k, v := range health {
		fmt.Fprintf(w, "%s: %s\n", k, v)
	}
}
