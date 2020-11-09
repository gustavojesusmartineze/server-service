//Package users has the users structure
//you can create a new user with NewUser func
//and set the an id for the user with SetID func
package users

import (
	"crypto/rand"
	"encoding/hex"
)

//User struct
//can be formated to JSON
//also omitempty allows us to handle zero values on
//FirstName, LastName and Username
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

//SetID Creates a random ID for the user using crypto/rand library
//uuid will work better
func (u *User) SetID() {
	buf := make([]byte, 16)
	_, err := rand.Read(buf)
	if err != nil {
		panic(err) // out of randomness, should never happen
	}
	u.ID = hex.EncodeToString(buf)
}
