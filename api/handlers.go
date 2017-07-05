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
	var response LoginResponse
	response.Success = true
	response.Token = token
	if err := json.NewEncoder(w).Encode(response); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}
}

type SimpleResponse struct {
	Success bool `json:"success"`
}

func register(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r);
	succes := model.RegisterUser(vars["login"], vars["pass"])

	w.WriteHeader(http.StatusCreated)
	var response SimpleResponse
	response.Success = succes
	if err := json.NewEncoder(w).Encode(response); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}
}

func getInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["id"] {
		card, err := model.GetInfoById(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
		}
		if err := json.NewEncoder(w).Encode(card); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
		}
	} else if vars["link"] {
		card, err := model.GetInfoByLink(vars["link"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
		}
		if err := json.NewEncoder(w).Encode(card); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
		}
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func createCard(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	id, err := model.ValidToken(vars["token"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}

	err = model.CreateCard(id)
	if err!= nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}
	 var response SimpleResponse
	response.Success = true

	if err := json.NewEncoder(w).Encode(response); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}
}

type UpdateRequest struct {
	Token string `json:"token"`
	Update map[string]string `json:"update"`
} 

func update(w http.ResponseWriter, r *http.Request)  {
	decoder := json.NewDecoder(r.Body);
	var vars UpdateRequest
	err := decoder.Decode(vars)
	id, err := model.ValidToken(vars.Token)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}

	err = model.Update(id, vars.Update);

}

func deleteCard(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	id, err := model.ValidToken(vars["token"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}

	err = model.DeleteUserInfo(id)
	if err!= nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}
	var response SimpleResponse
	response.Success = true

	if err := json.NewEncoder(w).Encode(response); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}
}