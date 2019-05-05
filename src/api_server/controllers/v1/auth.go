package controllers

import (
	"encoding/json"
	"net/http"

	common "../../common"
	"../../models"
)

// PublishToken controller
var PublishToken = func(w http.ResponseWriter, r *http.Request) {

	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user) //decode the request body into struct and failed if any error occur
	if err != nil {
		common.Respond(w, common.Message(false, "Invalid request"))
		return
	}

	resp := models.PublishToken(user.Name, user.Password)
	common.Respond(w, resp)
}
