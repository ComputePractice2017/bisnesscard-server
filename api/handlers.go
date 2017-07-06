package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ComputePractice2017/bisnesscard-server/model"
)

type LoginResponse struct {
	Success bool   `json:"success"`
	Token   string `json:"token"`
}

func login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var vars RegisterRequest
	err := decoder.Decode(&vars)
	token, err := model.Login(vars.Login, vars.Pass)
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

type RegisterRequest struct {
	Login string `json:"login"`
	Pass  string `json:"pass"`
}

func register(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var vars RegisterRequest
	err := decoder.Decode(&vars)

	succes, err := model.RegisterUser(vars.Login, vars.Pass)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}
	w.WriteHeader(http.StatusCreated)
	var response SimpleResponse
	response.Success = succes
	if err := json.NewEncoder(w).Encode(response); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}
}

type TokenIdReuest struct {
	Token string `json:"token"`
	Link  string `json:"link"`
	ID    string `json:"id"`
}

func getInfo(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var vars TokenIdReuest
	_ = decoder.Decode(&vars)
	if vars.ID != "" {
		card, err := model.GetInfoById(vars.ID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
		}
		if err := json.NewEncoder(w).Encode(card); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
		}
	} else if vars.Link != "" {
		card, err := model.GetInfoByLink(vars.Link)
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

func createCard(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var vars TokenIdReuest
	err := decoder.Decode(&vars)
	println(vars.Token)
	id, err := model.ValidToken(vars.Token)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}

	err = model.CreateCard(id)
	if err != nil {
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

type UpdateFields struct {
	Desc string
}

type UpdateRequest struct {
	Token  string          `json:"token"`
	Update model.User_info `json:"update"`
}

func update(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var vars UpdateRequest
	err := decoder.Decode(&vars)
	id, err := model.ValidToken(vars.Token)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}

	err = model.Update(id, vars.Update)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}

}

func deleteCard(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var vars TokenIdReuest
	err := decoder.Decode(&vars)
	id, err := model.ValidToken(vars.Token)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}

	err = model.DeleteUserInfo(id)
	if err != nil {
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
