package data

// Reaction represents a user's response to a photo
type Reaction struct {
	ID       string `json:"id"`
	Photo    *Photo `json:"photo"`
	Reaction int    `json:"reaction"`
	User     *User  `json:"user"`
}

// Photo is the binary image data and metadata
type Photo struct {
	ID   string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
}

// User is a application user, including 'Anonymous'
type User struct {
	ID     string   `json:"id"`
	Name   string   `json:"name"`
	Photos []*Photo `json:"photos"`
}
