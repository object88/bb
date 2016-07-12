package data

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Init bootstraps some cheap datas
func Init() {
	s := GetSession()
	dbName := "bbgraph"

	usersCollection := getCollection(s, dbName, "users")
	var users = []User{
		User{bson.ObjectIdHex("000000000000000000000000"), "Anonymous"},
		User{bson.ObjectIdHex("000000000000000000000001"), "Root"},
		User{bson.NewObjectId(), "Other"},
	}
	insertIntoCollection(usersCollection, users[0], users[1], users[2])

	photosCollection := getCollection(s, dbName, "photos")
	var photos = []Photo{
		Photo{bson.NewObjectId(), "What's-it", users[1].ID},
		Photo{bson.NewObjectId(), "Who's-it", users[1].ID},
		Photo{bson.NewObjectId(), "How's-it", users[1].ID},
	}
	insertIntoCollection(photosCollection, photos[0], photos[1], photos[2])

	reactionsCollection := getCollection(s, dbName, "reactions")
	var reactions = []Reaction{
		Reaction{bson.NewObjectId(), photos[0].ID, 1, users[2].ID},
	}
	insertIntoCollection(reactionsCollection, reactions[0])
}

func insertIntoCollection(c *mgo.Collection, docs ...interface{}) {
	c.RemoveAll(nil)
	e := c.Insert(docs...)
	if e != nil {
		fmt.Printf("Error inserting to collection '%s': %s\n", c.Name, e)
	}
}

func getCollection(s *mgo.Session, dbName string, collectionName string) *mgo.Collection {
	collection := s.DB(dbName).C(collectionName)
	return collection
}

// GetPhoto returns a *Photo with the matching id
func GetPhoto(id string) *Photo {
	oid := bson.ObjectIdHex(id)
	s := GetSession()
	c := getCollection(s, "bbgraph", "photos")
	query := c.Find(bson.M{"_id": oid})
	var p *Photo
	query.One(p)
	return p
}

// GetPhotos returns all photos
func GetPhotos() []*Photo {
	s := GetSession()
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
	oid := bson.ObjectIdHex(id)
	s := GetSession()
	c := getCollection(s, "bbgraph", "users")
	query := c.Find(bson.M{"_id": oid})
	var u *User
	query.One(u)
	return u
}

// GetViewer returns the current user
func GetViewer() *User {
	return GetUser("000000000000000000000001")
}

// PhotoSliceToPhotoPointerSlice turns the []Photo into []*Photo
func PhotoSliceToPhotoPointerSlice(photos []Photo) []*Photo {
	var photoPointerSlice = make([]*Photo, len(photos))
	for i, d := range photos {
		photoPointerSlice[i] = &d
	}
	return photoPointerSlice
}

// PhotosToInterfaceSlice gets an interface slice.
// See https://github.com/golang/go/wiki/InterfaceSlice
func PhotosToInterfaceSlice(photos ...*Photo) []interface{} {
	var interfaceSlice = make([]interface{}, len(photos))
	for i, d := range photos {
		interfaceSlice[i] = d
	}
	return interfaceSlice
}
