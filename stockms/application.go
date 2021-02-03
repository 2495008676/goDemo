package main

import (
	"fmt"
	"net/http"
	"stockms/controller"
	"stockms/framework"
	"time"
)

func main() {

	framework.InitDB()
	framework.CreateTable()

	server := &http.Server{
		Addr:        ":8080",
		Handler:     framework.Router,
		ReadTimeout: 5 * time.Second,
	}
	RegisterRouter(framework.Router)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("start server error")
	}
	fmt.Println("start server success")
}

func RegisterRouter(handler *framework.RouterHandler) {
	new(controller.UserController).Router(handler)
}
