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

// ModifyAndShowProject controller
var ModifyAndShowProject = func(w http.ResponseWriter, r *http.Request) {
	project := &models.Project{}
	vars := mux.Vars(r)

	projectID := vars["project_id"]
	existProject := ShowProject
	if existProject == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&project); err != nil {
		println(err)
		return
	}
	project.ModifyProject(projectID)
	resp := common.Message(true, "success")

	respProject := models.ShowProject(projectID)
	resp["project"] = respProject
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
