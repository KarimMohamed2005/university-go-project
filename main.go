package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)
func main() {
	r := mux.NewRouter().StrictSlash(true)

	initialMigration()
	r.HandleFunc("/operation", createOperation).Methods("POST")
	r.HandleFunc("/operations", getOperations).Methods("GET")
	r.HandleFunc("/status",getStatus).Methods("GET")
	r.PathPrefix("/files/").Handler(http.StripPrefix("/files/", http.FileServer(http.Dir("./files/"))))
	r.PathPrefix("/swagger-ui/").Handler(http.StripPrefix("/swagger-ui/", http.FileServer(http.Dir("./swagger-ui/"))))

	log.Fatal(http.ListenAndServe(":8080",r))
}

