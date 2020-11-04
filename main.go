package main

import (
	"github.com/gmartinez8/server/pkg/server"
)

func main() {
	s := server.NewServer(":4000")
	s.Handle("/", "GET", server.HandleHome)
	s.Handle("/home", "GET", server.HandleHome)
	s.Handle("/users", "GET", server.HandleUsers)
	s.Handle("/users", "POST", server.HandleCreateUsers)
	//You can get user by id /user/{id} method GET
	//Example: /user/5
	s.Handle("/user", "GET", server.HandleShowUser)
	//You can delete users by id /users/{id} method DELETE
	//Example: /users/5
	s.Handle("/users", "DELETE", server.HandleDeleteUsers)
	//You can edit users by id /users/{id}
	//Example: /users/5 method PUT
	s.Handle("/users", "PUT", server.HandleEditUsers)
	s.Run()
}
