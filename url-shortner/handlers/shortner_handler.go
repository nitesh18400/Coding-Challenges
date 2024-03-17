package handlers

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"url-shortner/db"
	"url-shortner/models"
)

func generateShortURLKey(longURL string, length int) string {
	// Calculate MD5 hash of the long URL
	hash := md5.Sum([]byte(longURL))
	hashString := hex.EncodeToString(hash[:])

	// Truncate hash to desired length
	if length > len(hashString) {
		length = len(hashString)
	}
	shortURLKey := hashString[:length]

	return shortURLKey
}

func ShortenUrl(w http.ResponseWriter, r *http.Request) {
	var url models.Url
	err := json.NewDecoder(r.Body).Decode(&url)

	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	url.ShortURL = generateShortURLKey(url.OriginalURL, 6)

	shortURL, err := db.ShortenUrl(url)
	if err != nil {
		http.Error(w, "Failed to shorten URL", http.StatusInternalServerError)
		return
	}

	fullShortURL := "http://" + r.Host + "/" + shortURL
	response := map[string]string{"short_url": fullShortURL}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
