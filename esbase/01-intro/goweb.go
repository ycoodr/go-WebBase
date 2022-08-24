package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Server on 3000 port")
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
