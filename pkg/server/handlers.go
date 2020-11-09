package server

import (
	"encoding/json"
	"net/http"
	"strings"
	"sync"

	"github.com/gmartinez8/server/pkg/users"
)

//ErrorResponse allows us to
type ErrorResponse struct {
	Message string
}

//Users of the system, will work better with a DB connection
//but for the task purpose not a requirement asked
//I'll test later with a DB connection
var usersdb = make(map[string]*users.User)
var mutex = &sync.Mutex{}

//HandleHome Handles Home Route and /
func HandleHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(" HandleHome Called \n"))
}

//HandleUsers list all users stored
func HandleUsers(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	response, err := json.Marshal(usersdb)
	if len(usersdb) == 0 {
		response, _ = json.Marshal(make([]users.User, 0))
	}
	mutex.Unlock()
	if err != nil {
		e, _ := json.Marshal(ErrorResponse{err.Error()})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(e)
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
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		e, _ := json.Marshal(ErrorResponse{err.Error()})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(e)
		return
	}
	user.SetID()
	response, err := json.Marshal(user)
	if err != nil {
		e, _ := json.Marshal(ErrorResponse{err.Error()})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(e)
		return
	}
	mutex.Lock()
	usersdb[user.ID] = &user
	mutex.Unlock()
	w.Write(response)
}

//HandleDeleteUsers delete user if exist
func HandleDeleteUsers(w http.ResponseWriter, r *http.Request) {
	pathElements := strings.Split(r.URL.Path, "/")
	index := len(pathElements)
	id := pathElements[index-1]
	mutex.Lock()
	_, ok := usersdb[id]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	delete(usersdb, id)
	mutex.Unlock()
	w.WriteHeader(http.StatusOK)
}

//HandleShowUser returns users data in JSON format
func HandleShowUser(w http.ResponseWriter, r *http.Request) {
	pathElements := strings.Split(r.URL.Path, "/")
	index := len(pathElements)
	id := pathElements[index-1]
	mutex.Lock()
	_, ok := usersdb[id]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	response, err := json.Marshal(usersdb[id])
	mutex.Unlock()
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		e, _ := json.Marshal(ErrorResponse{err.Error()})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(e)
		return
	}
	w.Write([]byte(response))
}

//HandleEditUsers delete user if exist
func HandleEditUsers(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user users.User
	//decodes the json from request and stores it in the value pointed by &user
	err := decoder.Decode(&user)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		e, _ := json.Marshal(ErrorResponse{err.Error()})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(e)
		return
	}
	pathElements := strings.Split(r.URL.Path, "/")
	index := len(pathElements)
	id := pathElements[index-1]
	//search for user with id received
	mutex.Lock()
	_, ok := usersdb[id]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if user.Email != "" {
		usersdb[id].Email = user.Email
	}
	if user.FirstName != "" {
		usersdb[id].FirstName = user.FirstName
	}
	if user.LastName != "" {
		usersdb[id].LastName = user.LastName
	}
	if user.Username != "" {
		usersdb[id].Username = user.Username
	}
	mutex.Unlock()
	response, err := json.Marshal(usersdb[id])
	if err != nil {
		e, _ := json.Marshal(ErrorResponse{err.Error()})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(e)
		return
	}
	w.Write([]byte(response))
}
