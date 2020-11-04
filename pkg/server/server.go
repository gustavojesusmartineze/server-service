package server

import (
	"fmt"
	"log"
	"net/http"
)

//Server struct allows us to define multiples servers if needed
type Server struct {
	port   string
	router *Router
}

//NewServer creates a new Server
func NewServer(port string) *Server {
	return &Server{
		port:   port,
		router: NewRouter(),
	}
}

//Run starts the server
func (s *Server) Run() error {
	fmt.Println("Server started on port 0.0.0.0", s.port)
	http.Handle("/", s.router)

	err := http.ListenAndServe(s.port, nil)
	if err != nil {
		log.Fatalln("Unable to run server on port", s.port)
		return err
	}

	log.Println("Server running on", s.port)
	return nil
}

//Handle defines/register the routes i want to handle
//also asign each route a HandlerFunc to handle it
//you can define each HandlerFunc in handlers.go file
func (s *Server) Handle(path string, method string, handler http.HandlerFunc) {
	//Check if the path already exists
	if !s.router.FindPath(path) {
		//If not path then create a new one
		s.router.rules[path] = make(map[string]http.HandlerFunc)
	}
	s.router.rules[path][method] = handler
}
