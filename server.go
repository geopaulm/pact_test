package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	StartServer()
}

func StartServer() {
	router := mux.NewRouter()
	router.HandleFunc("/users/{id:[0-9]+}", UserHandler)
	http.Handle("/", router)

	srv := &http.Server{
		Handler: router,
		Addr:    "0.0.0.0:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Started server. Listening..")
	log.Fatal(srv.ListenAndServe())
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := &User{
		Id:        vars["id"],
		FirstName: "Geo",
		LastName:  "Paul",
	}
	response, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
