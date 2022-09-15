package db

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var dsn = "yan:sunrise1@tcp(localhost:3306)/blog_db?charset=utf8mb4&parseTime=True&loc=Local"
var Database = func() (db *gorm.DB){
	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		fmt.Println("Error connection", err)
		panic(err)
	} else {
		fmt.Println("Connection is exist")
		return db
	}
}()

