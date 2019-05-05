package controllers

import (
	"encoding/json"
	"net/http"

	common "../../common"
	"../../models"
	"github.com/gorilla/mux"
)

// CreateUser controller
var CreateUser = func(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user) // リクエストのbody部を構造体にデコード
	if err != nil {
		common.Respond(w, common.Message(false, "Invalid request"))
		return
	}

	resp := user.CreateUser() // ユーザの作成
	common.Respond(w, resp)
}

// ListUser controller
var ListUser = func(w http.ResponseWriter, r *http.Request) {
	user := models.ListUser()
	resp := common.Message(true, "success")
	resp["users"] = user
	common.Respond(w, resp)
}

// ShowUser controller
var ShowUser = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID := params["user_id"]

	user := models.ShowUser(userID)
	resp := common.Message(true, "success")
	resp["user"] = user
	common.Respond(w, resp)
}

// ModifyAndShowUser controller
var ModifyAndShowUser = func(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	vars := mux.Vars(r)

	userID := vars["user_id"]
	existUser := ShowUser
	if existUser == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		println(err)
		return
	}
	user.ModifyUser(userID)
	resp := common.Message(true, "success")

	respUser := models.ShowUser(userID)
	resp["user"] = respUser
	common.Respond(w, resp)
}

// DeleteUser controller
var DeleteUser = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID := params["user_id"]

	models.DeleteUser(userID)
	resp := common.Message(true, "success")
	common.Respond(w, resp)
}
