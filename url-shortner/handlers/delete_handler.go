package handlers

import "net/http"
import "url-shortner/db"

func DeleteUrl(w http.ResponseWriter, r *http.Request) {
	// Extract the short URL from the request
	shortURL := r.URL.Path[1:]

	err := db.DeleteUrl(shortURL)

	if err != nil {
		http.Error(w, "Failed to delete Url", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("URL deleted successfully"))
}
