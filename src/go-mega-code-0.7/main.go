package main

import (
	"go-mega-code-0.7/controller"
	"go-mega-code-0.7/model"
	"net/http"

	"github.com/gorilla/context"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// Setup DB
	db := model.ConnectToDB()
	defer db.Close()
	model.SetDB(db)

	// Setup Controller
	controller.Startup()

	http.ListenAndServe(":8888", context.ClearHandler(http.DefaultServeMux))
}
