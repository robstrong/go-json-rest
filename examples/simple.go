package main

import (
	"github.com/ant0ine/go-json-rest"
	"net/http"
)

type User struct {
	Id   string
	Name string
}

func GetUser(w *rest.ResponseWriter, req *rest.Request) {
	user := User{
		Id:   req.PathParam("id"),
		Name: "Antoine",
	}
	w.WriteJSON(&user)
}

func main() {
	handler := rest.ResourceHandler{}
	handler.SetRoutes(
		rest.Route{"GET", "/users/:id", GetUser},
	)
	http.ListenAndServe(":8080", &handler)
}
