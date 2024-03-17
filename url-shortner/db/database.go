package db

import (
	"context"
	"log"
	"net/http"
	"time"
	"url-shortner/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func init() {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Could not connect to MongoDB:", err)
	}
	log.Println("Connected to MongoDB!")

	collection = client.Database("urlshortener").Collection("urls")

}

func ShortenUrl(url models.Url) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, url)
	if err != nil {
		return "", err
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()
	log.Println("Inserted new URL with ID:", id)
	return url.ShortURL, nil
}

func GetOriginalUrl(shortUrl string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var url models.Url
	err := collection.FindOne(ctx, bson.M{"short_url": shortUrl}).Decode(&url)
	if err != nil {
		log.Println("Error finding URL:", err)
		return "", err
	}
	log.Println("Found original URL:", url.OriginalURL)
	return url.OriginalURL, nil
}

func RedirectToOriginal(w http.ResponseWriter, r *http.Request, Shorturl string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var url models.Url
	err := collection.FindOne(ctx, bson.M{"shorturl": Shorturl}).Decode(&url)
	if err != nil {
		// Handle error (e.g., URL not found)
		return
	}
	http.Redirect(w, r, url.OriginalURL, http.StatusFound)

}

func DeleteUrl(shortUrl string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.DeleteOne(ctx, bson.M{"short_url": shortUrl})
	if err != nil {
		return err
	}
	return nil
}
