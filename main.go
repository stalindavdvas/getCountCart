package main

import (
	"getcountcart/database"
	"getcountcart/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors" // Middleware for CORS
)

func main() {
	// Initialize Redis connection
	client := database.InitRedis()
	defer client.Close()

	// Create a new router
	r := mux.NewRouter()

	// Define the route to get cart count
	r.HandleFunc("/api/cart/count", handlers.GetCartCount(client)).Methods("GET")

	// Configure CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://3.229.231.204:3000"}, // Allow requests from frontend
		AllowedMethods: []string{"GET"},
	})

	// Wrap the router with CORS middleware
	handler := corsHandler.Handler(r)

	// Start the server
	log.Println("Server started on http://localhost:8083")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
