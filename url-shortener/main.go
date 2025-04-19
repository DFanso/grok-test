package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

// URLStore holds the mapping of short to original URLs
type URLStore struct {
	store map[string]string
	mutex sync.RWMutex
}

func NewURLStore() *URLStore {
	return &URLStore{
		store: make(map[string]string),
	}
}

func (us *URLStore) Get(shortURL string) (string, bool) {
	us.mutex.RLock()
	defer us.mutex.RUnlock()
	url, exists := us.store[shortURL]
	return url, exists
}

func (us *URLStore) Put(shortURL, originalURL string) {
	us.mutex.Lock()
	defer us.mutex.Unlock()
	us.store[shortURL] = originalURL
}

func generateShortURL() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 6
	rand.Seed(time.Now().UnixNano())
	shortURL := make([]byte, length)
	for i := range shortURL {
		shortURL[i] = charset[rand.Intn(len(charset))]
	}
	return string(shortURL)
}

func shortenHandler(store *URLStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var request struct {
			URL string `json:"url"`
		}
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if request.URL == "" {
			http.Error(w, "URL is required", http.StatusBadRequest)
			return
		}

		shortURL := generateShortURL()
		store.Put(shortURL, request.URL)

		response := struct {
			ShortURL string `json:"short_url"`
		}{
			ShortURL: fmt.Sprintf("http://localhost:8080/%s", shortURL),
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func redirectHandler(store *URLStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		shortURL := r.URL.Path[1:] // Remove leading "/"
		if originalURL, exists := store.Get(shortURL); exists {
			http.Redirect(w, r, originalURL, http.StatusFound)
		} else {
			http.Error(w, "Short URL not found", http.StatusNotFound)
		}
	}
}

func main() {
	store := NewURLStore()
	http.HandleFunc("/shorten", shortenHandler(store))

	// Serve static files for the web interface
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Handle root path for web interface and redirects for short URLs
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.ServeFile(w, r, "static/index.html")
			return
		}
		redirectHandler(store)(w, r)
	})

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
