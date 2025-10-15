package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

// URL struct to store original and short URLs
type URL struct {
	ID           string    `json:"id"`
	OriginalURL  string    `json:"Original_Url"`
	ShortURL     string    `json:"Short_Url"`
	CreationDate time.Time `json:"Creation_date"`
}

// In-memory storage
var urlDB = make(map[string]URL)

// Generate 8-character short URL using MD5
func generateShortUrl(originalURL string) string {
	hasher := md5.New()
	hasher.Write([]byte(originalURL))
	data := hasher.Sum(nil)
	hash := hex.EncodeToString(data)
	return hash[:8]
}

// Create a new short URL and store in memory
func createUrl(originalURL string) string {
	shortURL := generateShortUrl(originalURL)
	urlDB[shortURL] = URL{
		ID:           shortURL,
		OriginalURL:  originalURL,
		ShortURL:     shortURL,
		CreationDate: time.Now(),
	}
	return shortURL
}

// Retrieve URL from memory by short ID
func getUrl(id string) (URL, error) {
	url, ok := urlDB[id]
	if !ok {
		return URL{}, errors.New("URL not found")
	}
	return url, nil
}

// Root handler
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Go URL Shortener!")
}

// Handler to create short URL
func shortUrlHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		URL string `json:"url"`
	}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	shortURL := createUrl(data.URL)

	response := map[string]string{
		"Short_url": fmt.Sprintf("http://localhost:3000/redirect/%s", shortURL),
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

// Handler to redirect to original URL
func redirectUrlHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/redirect/"):]
	url, err := getUrl(id)
	if err != nil {
		http.Error(w, "Short URL not found", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, url.OriginalURL, http.StatusFound)
}

// Main function
func main() {

	fmt.Println("Starting Go URL Shortener server at port 3000...")

	http.HandleFunc("/shortner", shortUrlHandler)
	http.HandleFunc("/shortner/", shortUrlHandler)
	http.HandleFunc("/redirect/", redirectUrlHandler)
	http.HandleFunc("/", handler)

	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

// func main() {
// 	fmt.Println("Starting Go URL Shortener server at port 3000...")

// 	http.HandleFunc("/shortner", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println(">>> HIT /shortner with", r.Method, "path:", r.URL.Path)
// 		shortUrlHandler(w, r)
// 	})

// 	http.HandleFunc("/shortner/", shortUrlHandler)
// 	http.HandleFunc("/redirect/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println(">>> HIT /redirect with", r.Method, "path:", r.URL.Path)
// 		redirectUrlHandler(w, r)
// 	})

// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println(">>> HIT / (fallback) with", r.Method, "path:", r.URL.Path)
// 		fmt.Fprintln(w, "Welcome to Go URL Shortener!")
// 	})

// 	if err := http.ListenAndServe(":3000", nil); err != nil {
// 		fmt.Println("Error starting server:", err)
// 	}
// }
