package models

import "gorm/db"

type User struct {

    Id int64 /* `yaml:"id"` `xml:"id"`*/`json:"id"`
    Username string /*`yaml:"username"` `xml:"username"`*/`json:"username"`
    Password string /*`yaml:"password"` //`xml:"password"`*/ `json:"password"`
    Email string /*`yaml:"email"` //`xml:"email"` */`json:"email"`
}

type Users []User

func MigrateUser(){
    db.Database.AutoMigrate(User{})
}
