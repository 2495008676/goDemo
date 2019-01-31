package main

import (
    "go-mega-code-0.5/controller"
    "go-mega-code-0.5/model"
    "net/http"

    _ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
    // Setup DB
    db := model.ConnectToDB()
    defer db.Close()
    model.SetDB(db)

    // Setup Controller
    controller.Startup()

    http.ListenAndServe(":8888", nil)
}
