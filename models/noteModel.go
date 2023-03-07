package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Note struct {
	ID        primitive.ObjectID `bson:"_id"`
	Text      string             `json:"text"`
	Title     string             `json:"title"`
	Create_at string             `json:"create_at"`
	Update_at time.Time          `json:"update_at"`
	Note_id   time.Time          `json:"note_id"`
}
