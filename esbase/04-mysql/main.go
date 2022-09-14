package main

import (
    "gomysql/db"
    "gomysql/models"
    "fmt"
    
)

func main() {
	db.Connect()
    // fmt.Println(db.ExistsTable("users"))
    // user := models.CreateUser("alex", "alex123", "alex@gmail.com")
    user := models.CreateUser("roel", "roel456", "roel@gmail.com")
    fmt.Println(user)
//    users := models.ListUsers()
//    fmt.Println(users)
    // user := models.GetUser(1)
    // fmt.Println(user)
    // user.Username = "juan"
    // user.Password = "juan789"
    // user.Email = "juan@gmail.com"
    // user.Save()
    // user.Delete()
   // db.TruncateTable("users")
    fmt.Println(models.ListUsers())
  //  db.CreateTable(models.UserSchema, "users")
     
    //db.Ping()
	db.Close()
    
}
