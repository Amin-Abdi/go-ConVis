package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func HandleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	//The url routes
	router.HandleFunc("/home", homePage).Methods("GET")
	router.HandleFunc("/operations", getRepresentation).Methods("GET")

	var address string = "127.0.0.1:8000"

	fmt.Println("Server is currently running on:", address)
	err := http.ListenAndServe(address, router)
	if err != nil {
		return
	}

}

//HomePage : Just for Testing purposes only
func homePage(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Endpoints called: HomePage")
	fmt.Println("HomePage function called")
}

func getRepresentation(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("getRepresentation called!")
	rw.Header().Set("Content-Type", "application/json")

	json.NewEncoder(rw).Encode(Operations)
}
