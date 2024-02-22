package main

import (
	"fmt"
	"log"
	"mongo_golang/configs"
	"mongo_golang/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	// run database
	configs.ConnectDb()

	// routes
	routes.UserRoute(router)

	fmt.Println("Listening to port 9000 \n")
	log.Fatal(http.ListenAndServe(":9000", router))
}
