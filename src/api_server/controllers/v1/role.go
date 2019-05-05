package controllers

import (
	"encoding/json"
	"net/http"

	common "../../common"
	"../../models"
	"github.com/gorilla/mux"
)

// CreateRole controller
var CreateRole = func(w http.ResponseWriter, r *http.Request) {
	role := &models.Role{}
	err := json.NewDecoder(r.Body).Decode(role) // リクエストのbody部を構造体にデコード
	if err != nil {
		common.Respond(w, common.Message(false, "Invalid request"))
		return
	}

	resp := role.CreateRole() // ユーザの作成
	common.Respond(w, resp)
}

// ListRole controller
var ListRole = func(w http.ResponseWriter, r *http.Request) {
	role := models.ListRole()
	resp := common.Message(true, "success")
	resp["roles"] = role
	common.Respond(w, resp)
}

// ShowRole controller
var ShowRole = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	roleID := params["role_id"]

	role := models.ShowRole(roleID)
	resp := common.Message(true, "success")
	resp["role"] = role
	common.Respond(w, resp)
}

// ModifyAndShowRole controller
var ModifyAndShowRole = func(w http.ResponseWriter, r *http.Request) {
	role := &models.Role{}
	vars := mux.Vars(r)

	roleID := vars["role_id"]
	existRole := ShowRole
	if existRole == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&role); err != nil {
		println(err)
		return
	}
	role.ModifyRole(roleID)
	resp := common.Message(true, "success")

	respRole := models.ShowRole(roleID)
	resp["role"] = respRole
	common.Respond(w, resp)
}

// DeleteRole controller
var DeleteRole = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	roleID := params["role_id"]

	models.DeleteRole(roleID)
	resp := common.Message(true, "success")
	common.Respond(w, resp)
}
