package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gmartinez8/server/pkg/users"
)

//Users of the system, will work better with a DB connection
//but for the task purpose not a requirement asked
//I'll test later with a DB connection
var usersdb []*users.User

//Handler type
type Handler func(w http.ResponseWriter, r *http.Request)

//HandleHome Handles Home Route and /
func HandleHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(" HandleHome Called \n"))
}

//HandleUsers list all users stored
func HandleUsers(w http.ResponseWriter, r *http.Request) {
	response, err := json.Marshal(usersdb)
	if len(usersdb) == 0 {
		response, _ = json.Marshal("No users registered yet")
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

//HandleCreateUsers Handles Home Route /
func HandleCreateUsers(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user users.User
	//decodes the json from request and stores it in the value pointed by &user
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}
	user.SetID()
	response, err := json.Marshal(user)
	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}
	usersdb = append(usersdb, &user)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

//HandleDeleteUsers delete user if exist
func HandleDeleteUsers(w http.ResponseWriter, r *http.Request) {
	pathElements := strings.Split(r.URL.Path, "/")
	index := len(pathElements)
	id := pathElements[index-1]
	for index, u := range usersdb {
		if u.ID == id {
			usersdb = append(usersdb[:index], usersdb[index+1:]...)
			response := "User width ID#" + u.ID + " Succesfully deleted"
			w.Write([]byte(response))
			return
		}
	}
	w.Write([]byte("User not found"))
}

//HandleShowUser returns users data in JSON format
func HandleShowUser(w http.ResponseWriter, r *http.Request) {
	pathElements := strings.Split(r.URL.Path, "/")
	index := len(pathElements)
	id := pathElements[index-1]
	for index, u := range usersdb {
		if u.ID == id {
			response, err := json.Marshal(usersdb[index])
			if err != nil {
				fmt.Fprintf(w, "error: %v", err)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(response))
			return
		}
	}
	w.Write([]byte("User not found"))
}

//HandleEditUsers delete user if exist
func HandleEditUsers(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user users.User
	//decodes the json from request and stores it in the value pointed by &user
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}
	pathElements := strings.Split(r.URL.Path, "/")
	index := len(pathElements)
	id := pathElements[index-1]
	//search for user with id received
	for index, u := range usersdb {
		if u.ID == id {
			if user.Email != "" {
				usersdb[index].Email = user.Email
			}
			if user.FirstName != "" {
				usersdb[index].FirstName = user.FirstName
			}
			if user.LastName != "" {
				usersdb[index].LastName = user.LastName
			}
			if user.Username != "" {
				usersdb[index].Username = user.Username
			}
			response, err := json.Marshal(usersdb[index])
			if err != nil {
				fmt.Fprintf(w, "error: %v", err)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(response))
			return
		}
	}
	w.Write([]byte("User not found"))
}
