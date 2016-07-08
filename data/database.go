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
	ID   string `json:"id"`
	Name string `json:"name"`
}

// User is a application user, including 'Anonymous'
type User struct {
	ID     string   `json:"id"`
	Name   string   `json:"name"`
	Photos []*Photo `json:"photos"`
}

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

// GetPhoto returns a *Photo with the matching id
func GetPhoto(id string) *Photo {
	for _, photo := range photos {
		if photo.ID == id {
			return photo
		}
	}
	return nil
}

// GetPhotos returns all photos
func GetPhotos() []*Photo {
	return photos
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

// PhotosToInterfaceSlice does some shit I don't understand the need for.
func PhotosToInterfaceSlice(photos ...*Photo) []interface{} {
	var interfaceSlice = make([]interface{}, len(photos))
	for i, d := range photos {
		interfaceSlice[i] = d
	}
	return interfaceSlice
}
