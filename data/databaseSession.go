package data

import (
	"fmt"
	"sync"

	"gopkg.in/mgo.v2"
)

var once sync.Once
var session *mgo.Session

// GetSession retreives the one and only MongoDB session
func GetSession() *mgo.Session {
	once.Do(func() {
		session = instantiate()
	})
	return session
}

func instantiate() *mgo.Session {
	s, err := mgo.Dial("127.0.0.1:27017/")

	if err != nil {
		fmt.Printf("Failed to create session to database: %s\n", err)
		panic(err)
	}

	fmt.Printf("Connected to mongo database.\n")
	return s
}
