package main

import (
	"log"
	"net/http"
	"todolist-service/handlers"

	"github.com/gorilla/mux"
)

func main() {
	handlerNodeChallenge := &handlers.NodeChallengeHandler{}
	handlerLeftRightEqual := &handlers.LeftRightEqualHandler{}

	router := mux.NewRouter()
	router.HandleFunc("/node-challenge", handlerNodeChallenge.NodeChallenge).Methods("GET")
	router.HandleFunc("/left-right-equal", handlerLeftRightEqual.LeftRightEqual).Methods("POST")

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8081", router))
}
