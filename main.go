package main

import (
	"fmt"
	"go-chart/app"
	"go-chart/controllers"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	//會員
	router.HandleFunc("/api/user/create", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")

	//圖表
	router.HandleFunc("/api/chartList/create", controllers.CreateChatList).Methods("POST")
	router.HandleFunc("/api/chartList/get", controllers.GetChatList).Methods("POST")

	router.HandleFunc("/api/chartData/create", controllers.CreateChatData).Methods("POST")
	router.HandleFunc("/api/chartData/get", controllers.GetChatData).Methods("POST")

	//圖片
	router.HandleFunc("/api/pic/creat", controllers.CreatePic).Methods("POST")
	router.HandleFunc("/api/pic/get", controllers.GetPic).Methods("POST")

	router.Use(app.JwtAuthentication) //jwt 解析

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}
