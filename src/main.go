package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	var router = mux.NewRouter()

	router.HandleFunc("/get", getFunction).Methods("GET")
	router.HandleFunc("/post", postFunction).Methods("POST")
	router.HandleFunc("/i/{urlparameter}", handleUrlMessage).Methods("GET")

	originCheck := handlers.AllowedOrigins([]string{"*"})
	headerCheck := handlers.AllowedHeaders([]string{"Authorization"})
	methodCheck := handlers.AllowedMethods([]string{"GET", "POST"})

	fmt.Println("Running server on PORT 3000!")
	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(originCheck, headerCheck, methodCheck)(router)))
}

func getFunction(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Working!!!")
}

func postFunction(w http.ResponseWriter, r *http.Request) {
	value1 := r.FormValue("key1")
	json.NewEncoder(w).Encode(map[string]string{"key1": value1})
}

func handleUrlMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	parameter := vars["urlparameter"]
	json.NewEncoder(w).Encode(map[string]string{"Parameter": parameter})
}
