package handlers

import (
	"log"
	"net/http"
	"url-shortner/db"
)

func RedirectToOriginal(w http.ResponseWriter, r *http.Request) {
	// Extract the short URL from the request
	shortURL := r.URL.Path[1:]
	log.Println("hulaka",shortURL)

	originalURL, err := db.GetOriginalUrl(shortURL)

	

	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, originalURL, http.StatusFound)
}
