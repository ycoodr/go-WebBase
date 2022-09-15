package main

import (
	"gorm/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	//"gorm/models"
)

func main(){
	//models.MigrateUser()

	mux := mux.NewRouter()

	mux.HandleFunc("/api/user/", handlers.GetUsers).Methods("GET")
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.GetUser).Methods("GET")
	mux.HandleFunc("/api/user/", handlers.CreateUser).Methods("POST")
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.UpdateUser).Methods("PUT")
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", mux))
}