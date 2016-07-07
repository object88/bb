package data

// Model structs
type User struct {
	Id     string   `json:"id"`
	Name   string   `json:"name"`
	Photos []*Photo `json:"photos"`
}

type Photo struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// Mock data
var viewer = &User{
	Id:   "1",
	Name: "Anonymous",
}
var photos = []*Photo{
	&Photo{"0", "What's-it"},
	&Photo{"1", "Who's-it"},
	&Photo{"2", "How's-it"},
}

// Data accessors
func GetPhoto(id string) *Photo {
	for _, photo := range photos {
		if photo.Id == id {
			return photo
		}
	}
	return nil
}
func GetPhotos() []*Photo {
	return photos
}
func GetUser(id string) *User {
	if id == viewer.Id {
		return viewer
	}
	return nil
}
func GetViewer() *User {
	return viewer
}
func PhotosToInterfaceSlice(photos ...*Photo) []interface{} {
	var interfaceSlice = make([]interface{}, len(photos))
	for i, d := range photos {
		interfaceSlice[i] = d
	}
	return interfaceSlice
}
