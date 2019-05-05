package controllers

import (
	"encoding/json"
	"net/http"

	common "../../common"
	"../../models"
	"github.com/gorilla/mux"
)

// CreateDashboardConfig controller
var CreateDashboardConfig = func(w http.ResponseWriter, r *http.Request) {
	dashboard := &models.DashboardConfig{}
	err := json.NewDecoder(r.Body).Decode(dashboard) // リクエストのbody部を構造体にデコード
	if err != nil {
		common.Respond(w, common.Message(false, "Invalid request"))
		return
	}

	resp := dashboard.CreateDashboardConfig() // ユーザの作成
	common.Respond(w, resp)
}

// ShowDashboardConfig controller
var ShowDashboardConfig = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	dashboardID := params["dashboard_id"]

	dashboard := models.ShowDashboardConfig(dashboardID)
	resp := common.Message(true, "success")
	resp["dashboard"] = dashboard
	common.Respond(w, resp)
}

// ModifyDashboardConfig controller
var ModifyDashboardConfig = func(w http.ResponseWriter, r *http.Request) {
	//params := mux.Vars(r)
	//dashboardID := params["dashboard_id"]

	dashboard := &models.DashboardConfig{}
	err := json.NewDecoder(r.Body).Decode(dashboard) // リクエストのbody部を構造体にデコード
	if err != nil {
		common.Respond(w, common.Message(false, "Invalid request"))
		return
	}

	//dashboard := models.ModifyDashboardConfig(dashboardID)
	resp := common.Message(true, "success")
	resp["dashboard"] = dashboard
	common.Respond(w, resp)
}

// DeleteDashboardConfig controller
var DeleteDashboardConfig = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	dashboardID := params["dashboard_id"]

	models.DeleteDashboardConfig(dashboardID)
	resp := common.Message(true, "success")
	common.Respond(w, resp)
}
