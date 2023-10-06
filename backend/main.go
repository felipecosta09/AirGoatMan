package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Serve static files from the relative path to your HTML file
	http.Handle("/", http.FileServer(http.Dir("http://frontend:80")))

	// Start the HTTP server on port 8080
	port := "8080"
	fmt.Printf("Server is running on port %s...\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
