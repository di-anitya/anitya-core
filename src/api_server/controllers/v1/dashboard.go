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

// ModifyAndShowDashboardConfig controller
var ModifyAndShowDashboardConfig = func(w http.ResponseWriter, r *http.Request) {
	dashboard := &models.DashboardConfig{}
	vars := mux.Vars(r)

	dashboardID := vars["dashboard_id"]
	existDashboard := ShowDashboardConfig
	if existDashboard == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&dashboard); err != nil {
		println(err)
		return
	}
	dashboard.ModifyDashboardConfig(dashboardID)
	resp := common.Message(true, "success")

	respDashboard := models.ShowDashboardConfig(dashboardID)
	resp["dashboard"] = respDashboard
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
