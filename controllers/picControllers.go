package controllers

import (
	"encoding/json"
	"go-chart/models"
	u "go-chart/utils"
	"net/http"
)

var CreatePic = func(w http.ResponseWriter, r *http.Request) {

	pic := &models.Pic{}

	err := json.NewDecoder(r.Body).Decode(pic)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := pic.Create()
	u.Respond(w, resp)
}

var GetPic = func(w http.ResponseWriter, r *http.Request) {

	data := models.Get()
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
