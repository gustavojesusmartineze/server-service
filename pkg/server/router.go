package server

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

//Router struct to handle multiples methods
type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

//NewRouter creates a new Router type
func NewRouter() *Router {
	return &Router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}

//checkPathRegex check if Path for a DELETE and PUT request
//are valid and have a number and not a word for this example
// url= /users/123 is valid
// url= /users/asd is not valid
func checkPathRegex(path string) (string, bool) {
	regexpPath := strings.Split(path, "/")
	index := len(regexpPath)
	match, err := regexp.MatchString("([0-9]+)", regexpPath[index-1])
	if err != nil {
		fmt.Println(err)
	}
	if match == true {
		path = strings.Replace(path, "/"+regexpPath[index-1], "", 1)
	}
	return path, match
}

//FindHandler finds the handler assigned to a route example GET /api
func (r *Router) FindHandler(path string, method string) (http.HandlerFunc, bool, bool) {
	_, pathExist := r.rules[path]
	if method == "DELETE" || method == "PUT" {
		path, pathExist = checkPathRegex(path)
	}
	handler, methodExist := r.rules[path][method]
	return handler, methodExist, pathExist
}

//FindPath verifies if path is valid
func (r *Router) FindPath(path string) bool {
	_, pathExist := r.rules[path]
	return pathExist
}

//ServeHTTP for HandlerFunc to check if our pathExist and methodExist
//this is executed on each request
func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	handler, methodExist, pathExist := r.FindHandler(request.URL.Path, request.Method)
	if !pathExist {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !methodExist {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	handler(w, request)
}
