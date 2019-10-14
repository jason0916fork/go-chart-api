package models

import (
	u "go-chart/utils"

	"github.com/jinzhu/gorm"
)

type Pic struct {
	gorm.Model
	Name string `json:"name"`
	Url  string `json:"Url"`
}

func (pic *Pic) Create() map[string]interface{} {
	GetDB().Create(pic)

	resp := u.Message(true, "success")
	resp["pic"] = pic
	return resp
}

func (pic *Pic) Get() map[string]interface{} {

	response := u.Message(true, "Account has been created")
	response["pic"] = pic
	return response
}
