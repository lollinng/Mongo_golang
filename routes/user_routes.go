package routes

import (
	"mongo_golang/controllers"

	"github.com/gorilla/mux"
)

func UserRoute(router *mux.Router) {

	router.HandleFunc("/user", controllers.CreateUser()).Methods("POST")
	router.HandleFunc("/user/{userId}", controllers.GetUser()).Methods("GET")
	router.HandleFunc("/user/{userId}", controllers.EditUser()).Methods("PUT")
	router.HandleFunc("/user/{userId}", controllers.DeleteUser()).Methods("DELETE")
	router.HandleFunc("/users", controllers.GetAllUser()).Methods("GET")
}
