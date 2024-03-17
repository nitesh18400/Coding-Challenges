package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Url struct {
    ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    OriginalURL string             `bson:"original_url" json:"original_url"`
    ShortURL    string             `bson:"short_url" json:"short_url"`
}

