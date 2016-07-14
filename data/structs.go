package data

import "gopkg.in/mgo.v2/bson"

// Reaction represents a user's response to a photo
type Reaction struct {
	ID       bson.ObjectId `json:"id"`
	Photo    bson.ObjectId `json:"photo"`
	Reaction int           `json:"reaction"`
	User     bson.ObjectId `json:"user"`
}

// Photo is the binary image data and metadata
type Photo struct {
	ID    bson.ObjectId `json:"id" bson:"_id"`
	Name  string        `json:"name" bson:"name"`
	Owner bson.ObjectId `json:"owner" bson:"owner"`
}

// User is a application user, including 'Anonymous'
type User struct {
	ID   bson.ObjectId `json:"id" bson:"_id"`
	Name string        `json:"name" bson:"name"`
}
