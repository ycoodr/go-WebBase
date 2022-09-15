package db

import (
	"database/sql"
	"fmt"
    _ "github.com/go-sql-driver/mysql"
)

const url = "yan:sunrise1@tcp(localhost:3306)/blog_db"

var db *sql.DB

func Connect() {
	connection, err := sql.Open("mysql", url)

	if err != nil {
		panic(err)
	}
	fmt.Println("Db connected")
	db = connection
}

func Close() {
	db.Close()
}

func Ping() {
    if err := db.Ping(); err != nil {
    
        panic(err)
    }
}

func ExistsTable(tablename string) bool{
    sql := fmt.Sprintf("show tables like '%s'", tablename)
	//rows, err := db.Query(sql)
	rows, err := Query(sql)
	if err != nil {
		fmt.Println("Error:", err)
	}

	return rows.Next()
}

func CreateTable(schema string, name string){
	if !ExistsTable(name){
		_, err := db.Exec(schema)
	 	if err != nil {
		fmt.Println(err)
	}
	}
	
}

func TruncateTable(tablename string){
	sql := fmt.Sprintf("Truncate %s", tablename)
	Exec(sql)
}

func Exec(query string, args ...interface{}) (sql.Result, error){
	Connect()
	result, err := db.Exec(query, args...)
	Close()
	if err != nil {
		fmt.Println(err)
	}
	return result, err
}

func Query(query string, args ...interface{}) (*sql.Rows, error){
	Connect()
	rows, err := db.Query(query, args...)
	Close()
	if err != nil {
		fmt.Println(err)
	}
	return rows, err
}
