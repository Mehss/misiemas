// main.go
package main

import (
	"log"
	"os"
)

func main() {
	log.Println("Starting the application...")

	// Initialize the Fiber app
	app := InitializeApp()

	// Determine port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	// Run the web service

	// Start the server
	log.Fatal(app.Listen("0.0.0.0:" + port))

}
