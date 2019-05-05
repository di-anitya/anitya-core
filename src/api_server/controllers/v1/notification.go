package controllers

import (
	"encoding/json"
	"net/http"

	common "../../common"
	"../../models"
	"github.com/gorilla/mux"
)

/////////////////////////////////////////////////////////

// CreateNotificationSlack controller
var CreateNotificationSlack = func(w http.ResponseWriter, r *http.Request) {
	notification := &models.NotificationSlack{}
	err := json.NewDecoder(r.Body).Decode(notification) // リクエストのbody部を構造体にデコード
	if err != nil {
		common.Respond(w, common.Message(false, "Invalid request"))
		return
	}

	resp := notification.CreateNotificationSlack() // ユーザの作成
	common.Respond(w, resp)
}

// ListNotificationSlack controller
var ListNotificationSlack = func(w http.ResponseWriter, r *http.Request) {
	notification := models.ListNotificationSlack()
	resp := common.Message(true, "success")
	resp["notifications"] = notification
	common.Respond(w, resp)
}

// ShowNotificationSlack controller
var ShowNotificationSlack = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	notificationID := params["notification_id"]

	notification := models.ShowNotificationSlack(notificationID)
	resp := common.Message(true, "success")
	resp["notification"] = notification
	common.Respond(w, resp)
}

// ModifyAndShowNotificationSlack controller
var ModifyAndShowNotificationSlack = func(w http.ResponseWriter, r *http.Request) {
	notificationSlack := &models.NotificationSlack{}
	vars := mux.Vars(r)

	notificationSlackID := vars["notificationSlack_id"]
	existNotificationSlack := ShowNotificationSlack
	if existNotificationSlack == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&notificationSlack); err != nil {
		println(err)
		return
	}
	notificationSlack.ModifyNotificationSlack(notificationSlackID)
	resp := common.Message(true, "success")

	respNotificationSlack := models.ShowNotificationSlack(notificationSlackID)
	resp["notificationSlack"] = respNotificationSlack
	common.Respond(w, resp)
}

// DeleteNotificationSlack controller
var DeleteNotificationSlack = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	notificationID := params["notification_id"]

	models.DeleteNotificationSlack(notificationID)
	resp := common.Message(true, "success")
	common.Respond(w, resp)
}

/////////////////////////////////////////////////////////

// CreateNotificationHistory controller
var CreateNotificationHistory = func(w http.ResponseWriter, r *http.Request) {
	notification := &models.NotificationHistory{}
	err := json.NewDecoder(r.Body).Decode(notification) // リクエストのbody部を構造体にデコード
	if err != nil {
		common.Respond(w, common.Message(false, "Invalid request"))
		return
	}

	resp := notification.CreateNotificationHistory() // ユーザの作成
	common.Respond(w, resp)
}

// ListNotificationHistory controller
var ListNotificationHistory = func(w http.ResponseWriter, r *http.Request) {
	notification := models.ListNotificationHistory()
	resp := common.Message(true, "success")
	resp["notifications"] = notification
	common.Respond(w, resp)
}

// ShowNotificationHistory controller
var ShowNotificationHistory = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	notificationID := params["notification_id"]

	notification := models.ShowNotificationHistory(notificationID)
	resp := common.Message(true, "success")
	resp["notification"] = notification
	common.Respond(w, resp)
}

// ModifyAndShowNotificationHistory controller
var ModifyAndShowNotificationHistory = func(w http.ResponseWriter, r *http.Request) {
	notificationHistory := &models.NotificationHistory{}
	vars := mux.Vars(r)

	notificationHistoryID := vars["notificationHistory_id"]
	existNotificationHistory := ShowNotificationHistory
	if existNotificationHistory == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&notificationHistory); err != nil {
		println(err)
		return
	}
	notificationHistory.ModifyNotificationHistory(notificationHistoryID)
	resp := common.Message(true, "success")

	respNotificationHistory := models.ShowNotificationHistory(notificationHistoryID)
	resp["notificationHistory"] = respNotificationHistory
	common.Respond(w, resp)
}

// DeleteNotificationHistory controller
var DeleteNotificationHistory = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	notificationID := params["notification_id"]

	models.DeleteNotificationHistory(notificationID)
	resp := common.Message(true, "success")
	common.Respond(w, resp)
}
