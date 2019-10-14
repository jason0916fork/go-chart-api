package controllers

import (
	"encoding/json"
	"go-chart/models"
	u "go-chart/utils"
	"net/http"
)

var CreateChatList = func(w http.ResponseWriter, r *http.Request) {

	chart := &models.ChartList{}

	err := json.NewDecoder(r.Body).Decode(chart)
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := chart.CreateList()
	u.Respond(w, resp)
}

var CreateChatData = func(w http.ResponseWriter, r *http.Request) {

	chart := &models.ChartData{}

	err := json.NewDecoder(r.Body).Decode(chart)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := chart.CreateData()
	u.Respond(w, resp)
}
var GetChatList = func(w http.ResponseWriter, r *http.Request) {

	data := models.GetList()
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}

var GetChatData = func(w http.ResponseWriter, r *http.Request) {

	data := models.GetData()
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
