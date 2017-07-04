package api

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"../model"

	"github.com/gorilla/mux"
)

type LoginResponse struct {
	Success bool `json:"success"`
	Token string `json:"token"`
}

func login(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r);
	token, err := model.Login(vars["login"], vars["pass"]);
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}
	w.WriteHeader(http.StatusCreated)
	var response LoginResponse;
	response.Success = true
	response.Token = token;
	if err := json.NewEncoder(w).Encode(response); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}
}

func register(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r);
	succes := model.RegisterUser(vars["login"], vars["pass"]);

	w.WriteHeader(http.StatusCreated)
	var response LoginResponse
	response.Success = succes
	if err := json.NewEncoder(w).Encode(response); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}
}

func getInfo(w http.ResponseWriter, r *http.Request) {

}

func createCard(w http.ResponseWriter, r *http.Request)  {

}

func update(w http.ResponseWriter, r *http.Request)  {

}

func delete(w http.ResponseWriter, r *http.Request)  {

}