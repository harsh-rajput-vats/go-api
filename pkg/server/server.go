package server

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Serve creates a new web server and starts listening
func Serve(addr string) {
	router := mux.NewRouter()

	// Server REST API routes
	serveAPI(router)

	// Create server instance with CORS middleware
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
		handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
	)(router)

	var httpServer = &http.Server{
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		Handler:      corsHandler,
		Addr:         addr,
	}

	// Start HTTP server
	log.Printf("Listening on %s\n", httpServer.Addr)
	log.Fatal(httpServer.ListenAndServe())
}
