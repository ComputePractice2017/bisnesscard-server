package api

import (
	"log"
	"net/http"
	"github.com/ComputePractice2017/bisnesscard-server/model"
	"github.com/gorilla/mux"
)

func Run()  {
	log.Println("Connecting to rethinkDB on localhost...")
	err := model.InitSession()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected")

	r := mux.NewRouter()
	r.HandleFunc("/login", login).Methods("POST")
	r.HandleFunc("/register", register).Methods("POST")

	r.HandleFunc("/getInfo", getInfo).Methods("GET")
	r.HandleFunc("/create", createCard).Methods("POST")
	r.HandleFunc("/update", update).Methods("PUT")
	r.HandleFunc("/delete", deleteCard).Methods("DELETE")

	log.Println("Running the server on port 8000...")
	http.ListenAndServe(":8000", r)

}
