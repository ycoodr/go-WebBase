package handlers

import (
	"net/http"
	// "fmt"
	// "apirest/db"
	"apirest/models"
	"encoding/json"
	//"encoding/xml"
	//"gopkg.in/yaml.v2"
	"strconv"
	"github.com/gorilla/mux"
)

func GetUsers(rw http.ResponseWriter, r *http.Request){
	// fmt.Fprintln(rw, "List of users")

	if users, err := models.ListUsers(); err != nil {
		models.SendNoFound(rw)
	} else {
		models.SendData(rw, users)
	}

	//rw.Header().Set("Content-Type", "application/json")
	//rw.Header().Set("Content-Type", "text/xml")
	//db.Connect()
	// users, _ := models.ListUsers()
	// db.Close()
	//output, _ := json.Marshal(users)
	//output, _ := xml.Marshal(users)
	//output, _ := yaml.Marshal(users)
	//fmt.Fprintln(rw, string(output))
}

func GetUser(rw http.ResponseWriter, r *http.Request){
	//fmt.Fprintln(rw, "user")
	// rw.Header().Set("Content-Type", "application/json")

	// vars := mux.Vars(r)
	// userId, _ := strconv.Atoi(vars["id"])
	// db.Connect()
	// user, _ := models.GetUser(userId)
	// db.Close()
	// output, _ := json.Marshal(user)
	// fmt.Fprintln(rw, string(output))
	//if user, err := models.GetUser()
	if user, err := GetUserByRequest(r); err != nil {
		models.SendNoFound(rw)
	} else {
		models.SendData(rw, user)
	}
}

func CreateUser(rw http.ResponseWriter, r *http.Request){
	//fmt.Fprintln(rw, "create a user")
	//rw.Header().Set("Content-Type", "application/json")

	user := models.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		//fmt.Fprintln(rw, http.StatusUnprocessableEntity)
		models.SendUnprocessableEntity(rw)
	} else {
		//db.Connect()
		user.Save()
		models.SendData(rw, user)
		//db.Close()
	}
	
	// output, _ := json.Marshal(user)
	// fmt.Fprintln(rw, string(output))
}

func UpdateUser(rw http.ResponseWriter, r *http.Request){
	//fmt.Fprintln(rw, "update a user")
	//rw.Header().Set("Content-Type", "application/json")

	//user := models.User{}
	var userId int64
	if user, err := GetUserByRequest(r); err != nil {
		models.SendNoFound(rw)
	} else {
		userId = user.Id
	}

	user := models.User{}
	decoder := json.NewDecoder(r.Body)

	// if err := decoder.Decode(&user); err != nil {
	// 	fmt.Fprintln(rw, http.StatusUnprocessableEntity)
	// } else {
	// 	db.Connect()
	// 	user.Save()
	// 	db.Close()
	// }

	if err := decoder.Decode(&user); err != nil {
		models.SendUnprocessableEntity(rw)
	} else {
		user.Id = userId
		user.Save()
		models.SendData(rw, user)
	}
	
	// output, _ := json.Marshal(user)
	// fmt.Fprintln(rw, string(output))
}

func DeleteUser(rw http.ResponseWriter, r *http.Request){
	//fmt.Fprintln(rw, "delete a user")
	// rw.Header().Set("Content-Type", "application/json")

	// vars := mux.Vars(r)
	// userId, _ := strconv.Atoi(vars["id"])
	// db.Connect()
	// user, _ := models.GetUser(userId)
	// user.Delete()
	// db.Close()
	// output, _ := json.Marshal(user)
	// fmt.Fprintln(rw, string(output))

	if user, err := GetUserByRequest(r); err != nil {
		models.SendNoFound(rw)
	} else {
		user.Delete()
		models.SendData(rw, user)
	}
}

func GetUserByRequest(r *http.Request)(models.User, error){
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])

	if user, err := models.GetUser(userId); err != nil {
		return *user, err
	} else {
		return *user, nil
	}
}