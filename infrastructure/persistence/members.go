package persistence

import "go.mongodb.org/mongo-driver/bson/primitive"

type Member struct {
	ID      string               `json:"id" bson:"_id"`
	Name    string               `json:"name"`
	MemoIDs []primitive.ObjectID `json:"memoIds"`
}
