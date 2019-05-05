package controllers

import (
	"encoding/json"
	"net/http"

	common "../../common"
	"../../models"
	"github.com/gorilla/mux"
)

///////////////////////////////////////////////////////////

// CreateHTTPMonitoringConfig controller
var CreateHTTPMonitoringConfig = func(w http.ResponseWriter, r *http.Request) {
	monitoring := &models.HTTPMonitoringConfig{}
	err := json.NewDecoder(r.Body).Decode(monitoring) // リクエストのbody部を構造体にデコード
	if err != nil {
		common.Respond(w, common.Message(false, "Invalid request"))
		return
	}

	resp := monitoring.CreateHTTPMonitoringConfig() // ユーザの作成
	common.Respond(w, resp)
}

// ListHTTPMonitoringConfig controller
var ListHTTPMonitoringConfig = func(w http.ResponseWriter, r *http.Request) {
	monitoring := models.ListHTTPMonitoringConfig()
	resp := common.Message(true, "success")
	resp["monitorings"] = monitoring
	common.Respond(w, resp)
}

// ShowHTTPMonitoringConfig controller
var ShowHTTPMonitoringConfig = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	monitoringID := params["monitoring_id"]

	monitoring := models.ShowHTTPMonitoringConfig(monitoringID)
	resp := common.Message(true, "success")
	resp["monitoring"] = monitoring
	common.Respond(w, resp)
}

// ModifyHTTPMonitoringConfig controller
var ModifyHTTPMonitoringConfig = func(w http.ResponseWriter, r *http.Request) {
	//params := mux.Vars(r)
	//monitoringID := params["monitoring_id"]

	monitoring := &models.HTTPMonitoringConfig{}
	err := json.NewDecoder(r.Body).Decode(monitoring) // リクエストのbody部を構造体にデコード
	if err != nil {
		common.Respond(w, common.Message(false, "Invalid request"))
		return
	}

	//monitoring := models.ModifyHTTPMonitoringConfig(monitoringID)
	resp := common.Message(true, "success")
	resp["monitoring"] = monitoring
	common.Respond(w, resp)
}

// DeleteHTTPMonitoringConfig controller
var DeleteHTTPMonitoringConfig = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	monitoringID := params["monitoring_id"]

	models.DeleteHTTPMonitoringConfig(monitoringID)
	resp := common.Message(true, "success")
	common.Respond(w, resp)
}

///////////////////////////////////////////////////////////

// CreateDNSMonitoringConfig controller
var CreateDNSMonitoringConfig = func(w http.ResponseWriter, r *http.Request) {
	monitoring := &models.DNSMonitoringConfig{}
	err := json.NewDecoder(r.Body).Decode(monitoring) // リクエストのbody部を構造体にデコード
	if err != nil {
		common.Respond(w, common.Message(false, "Invalid request"))
		return
	}

	resp := monitoring.CreateDNSMonitoringConfig() // ユーザの作成
	common.Respond(w, resp)
}

// ListDNSMonitoringConfig controller
var ListDNSMonitoringConfig = func(w http.ResponseWriter, r *http.Request) {
	monitoring := models.ListDNSMonitoringConfig()
	resp := common.Message(true, "success")
	resp["monitorings"] = monitoring
	common.Respond(w, resp)
}

// ShowDNSMonitoringConfig controller
var ShowDNSMonitoringConfig = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	monitoringID := params["monitoring_id"]

	monitoring := models.ShowDNSMonitoringConfig(monitoringID)
	resp := common.Message(true, "success")
	resp["monitoring"] = monitoring
	common.Respond(w, resp)
}

// ModifyDNSMonitoringConfig controller
var ModifyDNSMonitoringConfig = func(w http.ResponseWriter, r *http.Request) {
	//params := mux.Vars(r)
	//monitoringID := params["monitoring_id"]

	monitoring := &models.HTTPMonitoringConfig{}
	err := json.NewDecoder(r.Body).Decode(monitoring) // リクエストのbody部を構造体にデコード
	if err != nil {
		common.Respond(w, common.Message(false, "Invalid request"))
		return
	}

	//monitoring := models.ModifyDNSMonitoringConfig(monitoringID)
	resp := common.Message(true, "success")
	resp["monitoring"] = monitoring
	common.Respond(w, resp)
}

// DeleteDNSMonitoringConfig controller
var DeleteDNSMonitoringConfig = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	monitoringID := params["monitoring_id"]

	models.DeleteDNSMonitoringConfig(monitoringID)
	resp := common.Message(true, "success")
	common.Respond(w, resp)
}