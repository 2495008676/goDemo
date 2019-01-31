package model

import (
    "go-mega-code-1.3/config"
    "log"

    "github.com/jinzhu/gorm"
)

var db *gorm.DB

// SetDB func
func SetDB(database *gorm.DB) {
    db = database
}

// ConnectToDB func
func ConnectToDB() *gorm.DB {
    if db != nil {
        return db
    }
    connectingStr := config.GetMysqlConnectingString()
    log.Println("Connet to db...")
    db, err := gorm.Open("mysql", connectingStr)
    if err != nil {
        panic("Failed to connect database")
    }
    db.SingularTable(true)
    return db
}
