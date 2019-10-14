package models

import (
	u "go-chart/utils"

	"github.com/jinzhu/gorm"
)

type ChartList struct {
	gorm.Model
	Name  string `json:"name"`
	XName string `json:"XName"`
	YName string `json:"YName"`
}

type ChartData struct {
	gorm.Model
	Name  string `json:"name"`
	XData int    `json:"XData"`
	YData int    `json:"YData"`
}

func (chartList *ChartList) CreateList() map[string]interface{} {

	GetDB().Create(chartList)

	resp := u.Message(true, "success")
	resp["chartList"] = chartList
	return resp
}

func (chartdata *ChartData) CreateData() map[string]interface{} {

	GetDB().Create(chartdata)

	resp := u.Message(true, "success")
	resp["chartdata"] = chartdata
	return resp
}

func GetList() *ChartList {

	chartlist := &ChartList{}
	err := GetDB().Table("chartlist").Error
	if err != nil {
		return nil
	}
	return chartlist
}

func GetData() *ChartData {
	chartdata := &ChartData{}
	return chartdata
	//TODO-fix
}
