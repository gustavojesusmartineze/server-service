package main

import (
	"../service-task/pkg/server"
)

func main() {
	s := server.NewServer(":4000")
	s.Handle("/", "GET", server.HandleHome)
	s.Handle("/home", "GET", server.HandleHome)
	s.Handle("/users", "GET", server.HandleUsers)
	s.Handle("/users", "POST", server.HandleCreateUsers)
	s.Handle("/users", "DELETE", server.HandleDeleteUsers)
	s.Handle("/users", "PUT", server.HandleEditUsers)
	s.Run()
}
