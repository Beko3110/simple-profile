package main

import (
	"log"
	"net/http"
	"os"
)

// homePage renders the home HTML page
func homePage(w http.ResponseWriter, r *http.Request) {
	serveFile(w, r, "static/home.html")
}

// coursePage renders the course HTML page
func coursePage(w http.ResponseWriter, r *http.Request) {
	serveFile(w, r, "static/courses.html")
}

// aboutPage renders the about HTML page
func aboutPage(w http.ResponseWriter, r *http.Request) {
	serveFile(w, r, "static/about.html")
}

// contactPage renders the contact HTML page
func contactPage(w http.ResponseWriter, r *http.Request) {
	serveFile(w, r, "static/contact.html")
}

// serveFile is a helper function to handle errors when serving files
func serveFile(w http.ResponseWriter, r *http.Request, filepath string) {
	// Check if file exists
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		http.NotFound(w, r)
		return
	}
	// Serve the file
	http.ServeFile(w, r, filepath)
}

func main() {
	// Define handlers
	http.HandleFunc("/home", homePage)
	http.HandleFunc("/courses", coursePage)
	http.HandleFunc("/about", aboutPage)
	http.HandleFunc("/contact", contactPage)

	// Start the server with a dynamic address (better for flexibility)
	port := "8080"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}

	log.Printf("Starting server on port %s...", port)
	err := http.ListenAndServe("0.0.0.0:"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
