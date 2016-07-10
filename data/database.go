package data

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Mock data
var photos = []*Photo{
	&Photo{"0", "What's-it"},
	&Photo{"1", "Who's-it"},
	&Photo{"2", "How's-it"},
}
var users = []*User{
	&User{"0", "Anonymous", []*Photo{}},
	&User{"1", "Root", photos},
	&User{"2", "Other", []*Photo{}},
}
var reactions = []*Reaction{
	&Reaction{"0", photos[0], 1, users[2]},
}
var viewer = users[0]

func getSession() *mgo.Session {
	s, err := mgo.Dial("127.0.0.1:27017/")

	if err != nil {
		fmt.Printf("Failed to create session to database: %s\n", err)
		panic(err)
	}

	fmt.Printf("Connected to mongo database.\n")
	return s
}

func getCollection(s *mgo.Session, dbName string, collectionName string) *mgo.Collection {
	collection := s.DB(dbName).C(collectionName)
	return collection
}

// GetPhoto returns a *Photo with the matching id
func GetPhoto(id string) *Photo {
	oid := bson.ObjectIdHex(id)
	s := getSession()
	c := getCollection(s, "bbgraph", "photos")
	query := c.Find(bson.M{"_id": oid})
	var p *Photo
	query.One(p)
	return p
}

// GetPhotos returns all photos
func GetPhotos() []*Photo {
	s := getSession()
	c := getCollection(s, "bbgraph", "photos")
	query := c.Find(nil)
	var photos []Photo
	query.All(&photos)
	if photos != nil {
		return PhotoSliceToPhotoPointerSlice(photos)
	}
	return []*Photo{}
}

// GetUser returns a *User with the matching id
func GetUser(id string) *User {
	if id == viewer.ID {
		return viewer
	}
	return nil
}

// GetViewer returns the current user
func GetViewer() *User {
	return viewer
}

// PhotoSliceToPhotoPointerSlice turns the []Photo into []*Photo
func PhotoSliceToPhotoPointerSlice(photos []Photo) []*Photo {
	var photoPointerSlice = make([]*Photo, len(photos))
	for i, d := range photos {
		photoPointerSlice[i] = &d
	}
	return photoPointerSlice
}

// PhotosToInterfaceSlice does some shit I don't understand the need for.
func PhotosToInterfaceSlice(photos ...*Photo) []interface{} {
	var interfaceSlice = make([]interface{}, len(photos))
	for i, d := range photos {
		interfaceSlice[i] = d
	}
	return interfaceSlice
}
