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
	handlerPieFireDire := &handlers.PieFireDireHandler{}

	router := mux.NewRouter()
	router.HandleFunc("/node-challenge", handlerNodeChallenge.NodeChallenge).Methods("GET")
	router.HandleFunc("/left-right-equal", handlerLeftRightEqual.LeftRightEqual).Methods("POST")
	router.HandleFunc("/beef/summary", handlerPieFireDire.PieFireDire).Methods("GET")

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
