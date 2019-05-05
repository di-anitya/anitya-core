package controllers

import (
	"encoding/json"
	"net/http"

	common "../../common"
	"../../models"
	"github.com/gorilla/mux"
)

// CreateProject controller
var CreateProject = func(w http.ResponseWriter, r *http.Request) {
	project := &models.Project{}
	err := json.NewDecoder(r.Body).Decode(project) // リクエストのbody部を構造体にデコード
	if err != nil {
		common.Respond(w, common.Message(false, "Invalid request"))
		return
	}

	resp := project.CreateProject() // ユーザの作成
	common.Respond(w, resp)
}

// ListProject controller
var ListProject = func(w http.ResponseWriter, r *http.Request) {
	project := models.ListProject()
	resp := common.Message(true, "success")
	resp["projects"] = project
	common.Respond(w, resp)
}

// ShowProject controller
var ShowProject = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	projectID := params["project_id"]

	project := models.ShowProject(projectID)
	resp := common.Message(true, "success")
	resp["project"] = project
	common.Respond(w, resp)
}

// ModifyProject controller
var ModifyProject = func(w http.ResponseWriter, r *http.Request) {
	//params := mux.Vars(r)
	//projectID := params["project_id"]

	project := &models.Project{}
	err := json.NewDecoder(r.Body).Decode(project) // リクエストのbody部を構造体にデコード
	if err != nil {
		common.Respond(w, common.Message(false, "Invalid request"))
		return
	}

	//project := models.ModifyProject(projectID)
	resp := common.Message(true, "success")
	resp["project"] = project
	common.Respond(w, resp)
}

// DeleteProject controller
var DeleteProject = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	projectID := params["project_id"]

	models.DeleteProject(projectID)
	resp := common.Message(true, "success")
	common.Respond(w, resp)
}
