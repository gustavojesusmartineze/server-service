package server

import (
	"encoding/json"
	"math/rand"
	"strconv"
	"time"
)

//User struct
type User struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Username  string `json:"username,omitempty"`
}

//NewUser creates a new user
func NewUser(id, Email, FirstName, LastName, Username string) *User {
	return &User{
		ID:        id,
		Email:     Email,
		FirstName: FirstName,
		LastName:  LastName,
		Username:  Username,
	}
}

//ToJSON method to avoid json.Marhsal() every where
func (u *User) ToJSON() ([]byte, error) {
	return json.Marshal(u)
}

//Creates a random ID for the user
//uuid will work better
func (u *User) setID() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	u.ID = strconv.Itoa(r1.Intn(1000))
}
