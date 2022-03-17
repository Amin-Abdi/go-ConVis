package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func startServer() {
	router := mux.NewRouter().StrictSlash(true)
	address := "127.0.0.1:8000"

	//The API endpoints urls
	router.HandleFunc("/home", homePage).Methods("GET")
	router.HandleFunc("/operations", getRepresentation).Methods("GET")

	fmt.Println("Server is currently running on:", address)
	err := http.ListenAndServe(address, router)
	if err != nil {
		return
	}

}

//Just for Testing purposes only
func homePage(rw http.ResponseWriter, r *http.Request) {
	_ = r
	fmt.Fprintf(rw, "Endpoints called: HomePage")
}

//get the representation list in json format
func getRepresentation(rw http.ResponseWriter, r *http.Request) {
	//fmt.Println("getRepresentation called!")
	rw.Header().Set("Content-Type", "application/json")
	//To enable Cors policy
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	_ = r
	err := json.NewEncoder(rw).Encode(Operations)
	if err != nil {
		return
	}
}
