package controllers

import (
	"encoding/json"
	"net/http"

	common "../../common"
	"../../models"
	"github.com/gorilla/mux"
)

// CreateProbePoint controller
var CreateProbePoint = func(w http.ResponseWriter, r *http.Request) {
	probepoint := &models.ProbePoint{}
	err := json.NewDecoder(r.Body).Decode(probepoint) // リクエストのbody部を構造体にデコード
	if err != nil {
		common.Respond(w, common.Message(false, "Invalid request"))
		return
	}

	resp := probepoint.CreateProbePoint() // ユーザの作成
	common.Respond(w, resp)
}

// ListProbePoint controller
var ListProbePoint = func(w http.ResponseWriter, r *http.Request) {
	probepoint := models.ListProbePoint()
	resp := common.Message(true, "success")
	resp["probepoints"] = probepoint
	common.Respond(w, resp)
}

// ShowProbePoint controller
var ShowProbePoint = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	probepointID := params["probepoint_id"]

	probepoint := models.ShowProbePoint(probepointID)
	resp := common.Message(true, "success")
	resp["probepoint"] = probepoint
	common.Respond(w, resp)
}

// ModifyAndShowProbePoint controller
var ModifyAndShowProbePoint = func(w http.ResponseWriter, r *http.Request) {
	probepoint := &models.ProbePoint{}
	vars := mux.Vars(r)

	probepointID := vars["probepoint_id"]
	existProbePoint := ShowProbePoint
	if existProbePoint == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&probepoint); err != nil {
		println(err)
		return
	}
	probepoint.ModifyProbePoint(probepointID)
	resp := common.Message(true, "success")

	respProbePoint := models.ShowProbePoint(probepointID)
	resp["probepoint"] = respProbePoint
	common.Respond(w, resp)
}

// DeleteProbePoint controller
var DeleteProbePoint = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	probepointID := params["probepoint_id"]

	models.DeleteProbePoint(probepointID)
	resp := common.Message(true, "success")
	common.Respond(w, resp)
}
