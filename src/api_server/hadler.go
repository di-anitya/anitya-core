package apiserver

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"../models"

	"github.com/gorilla/mux"
)

// Status is ..
func Status(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// UserIndex is ..
func UserIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(users); err != nil {
		panic(err)
	}
}

// UserShow is ..
func UserShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userID"]
	fmt.Fprintln(w, "User show:", userID)
}

// UserCreate is ..
func UserCreate(w http.ResponseWriter, r *http.Request) {
	var user models.User
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &user); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := RepoCreateUser(user)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}
