package data

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Init bootstraps some cheap datas
func Init() {

	var users = []User{
		User{bson.ObjectIdHex("000000000000000000000000"), "Anonymous"},
		User{bson.ObjectIdHex("000000000000000000000001"), "Root"},
		User{bson.NewObjectId(), "Other"},
		User{bson.NewObjectId(), "Nother"},
		User{bson.NewObjectId(), "Xother"},
	}
	clearCollection("users")
	for _, v := range users {
		insertIntoCollection("users", v)
	}

	var photos = []Photo{
		Photo{bson.NewObjectId(), "A", users[1].ID},
		Photo{bson.NewObjectId(), "B", users[1].ID},
		Photo{bson.NewObjectId(), "C", users[1].ID},
		Photo{bson.NewObjectId(), "D", users[1].ID},
		Photo{bson.NewObjectId(), "E", users[1].ID},
		Photo{bson.NewObjectId(), "F", users[1].ID},
		Photo{bson.NewObjectId(), "G", users[1].ID},
		Photo{bson.NewObjectId(), "H", users[1].ID},
		Photo{bson.NewObjectId(), "I", users[1].ID},
		Photo{bson.NewObjectId(), "J", users[1].ID},
		Photo{bson.NewObjectId(), "K", users[2].ID},
		Photo{bson.NewObjectId(), "L", users[2].ID},
		Photo{bson.NewObjectId(), "M", users[2].ID},
		Photo{bson.NewObjectId(), "N", users[2].ID},
		Photo{bson.NewObjectId(), "O", users[2].ID},
		Photo{bson.NewObjectId(), "P", users[2].ID},
		Photo{bson.NewObjectId(), "Q", users[3].ID},
		Photo{bson.NewObjectId(), "R", users[3].ID},
		Photo{bson.NewObjectId(), "S", users[3].ID},
		Photo{bson.NewObjectId(), "T", users[3].ID},
		Photo{bson.NewObjectId(), "U", users[3].ID},
		Photo{bson.NewObjectId(), "V", users[3].ID},
		Photo{bson.NewObjectId(), "W", users[3].ID},
		Photo{bson.NewObjectId(), "X", users[3].ID},
		Photo{bson.NewObjectId(), "Y", users[3].ID},
		Photo{bson.NewObjectId(), "Z", users[3].ID},
	}
	clearCollection("photos")
	for _, v := range photos {
		insertIntoCollection("photos", v)
	}

	var reactions = []Reaction{
		Reaction{bson.NewObjectId(), photos[0].ID, 1, users[2].ID},
		Reaction{bson.NewObjectId(), photos[0].ID, 2, users[3].ID},
		Reaction{bson.NewObjectId(), photos[0].ID, 1, users[4].ID},
		Reaction{bson.NewObjectId(), photos[1].ID, -1, users[2].ID},
		Reaction{bson.NewObjectId(), photos[2].ID, 1, users[2].ID},
		Reaction{bson.NewObjectId(), photos[2].ID, 2, users[3].ID},
		Reaction{bson.NewObjectId(), photos[2].ID, -1, users[4].ID},
	}
	clearCollection("reactions")
	for _, v := range reactions {
		insertIntoCollection("reactions", v)
	}
}

func clearCollection(name string) {
	s := GetSession()
	dbName := "bbgraph"
	c := getCollection(s, dbName, name)

	c.RemoveAll(nil)
}

func insertIntoCollection(name string, doc interface{}) {
	s := GetSession()
	dbName := "bbgraph"
	c := getCollection(s, dbName, name)

	e := c.Insert(doc)
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

func GetPhotoReaction(oid bson.ObjectId) float64 {
	s := GetSession()
	c := getCollection(s, "bbgraph", "reactions")
	match := bson.M{"$match": bson.M{"photo": oid}}
	group := bson.M{
		"$group": bson.M{
			"_id": nil,
			"avg": bson.M{"$avg": "$reaction"},
		},
	}
	pipe := c.Pipe([]bson.M{match, group})
	var result bson.M
	err := pipe.One(&result)
	if err != nil {
		return 0
		//handle error
	}
	average := result["avg"].(float64)
	return average
}

// GetPhotos returns all photos
func GetPhotos() []*Photo {
	s := GetSession()
	c := getCollection(s, "bbgraph", "photos")
	query := c.Find(nil)
	var photos []*Photo
	query.All(&photos)
	if photos != nil {
		return photos
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
