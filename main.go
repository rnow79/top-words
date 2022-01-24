// Welcome to "Top Words" project. This project has two parts:
//  - A service that accepts input as text, n as int, and provides a json
//    with the n top used words, and times of occurence.
//  - A tiny frontend for testing the service.
package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// HTTP port to use.
const port int = 80

// HTTP server address (leave blank for default ip).
const serverAddr string = ""

// Main function.
func main() {
	// Create router.
	router := mux.NewRouter()
	// Register static html (for the frontend).
	router.PathPrefix("/front").Handler(http.StripPrefix("/front", http.FileServer(http.Dir("./front/"))))
	// Register default redirection.
	router.HandleFunc("/", redirectToFront).Methods(http.MethodGet)
	// Register our service.
	router.HandleFunc("/top", postTopWords).Methods(http.MethodPost)
	// Declare and start the web server.
	server := &http.Server{Handler: router, Addr: serverAddr + ":" + strconv.Itoa(port), WriteTimeout: 10 * time.Second, ReadTimeout: 10 * time.Second}
	server.ListenAndServe()
}
