package main

import (
	"go-mega-code-1.5/controller"
	"go-mega-code-1.5/model"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/context"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	// Setup DB
	db := model.ConnectToDB()
	defer db.Close()
	model.SetDB(db)

	// Setup Controller
	controller.Startup()

	port := os.Getenv("PORT")
	log.Println("Running on port: ", port)
	http.ListenAndServe(":"+port, context.ClearHandler(http.DefaultServeMux))
}
