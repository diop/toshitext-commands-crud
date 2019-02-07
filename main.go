package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Commands!")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", helloWorld).Methods("GET")
	myRouter.HandleFunc("/commands", AllCommands).Methods("GET")
	myRouter.HandleFunc("/command/{name}/{instructions}", NewCommand).Methods("POST")
	myRouter.HandleFunc("/command/{name}", DeleteCommand).Methods("DELETE")
	myRouter.HandleFunc("/command/{name}/{instructions}", UpdateCommand).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	fmt.Println("GORM Sandbox")
	InitialMigration()
	handleRequests()
}
