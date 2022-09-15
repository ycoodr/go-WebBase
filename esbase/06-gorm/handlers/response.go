package handlers

import (
	
	"net/http"
	"encoding/json"
	"fmt"
)

func sendData(rw http.ResponseWriter, data interface{}, status int){
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(status)

	output, _ := json.Marshal(&data)
	fmt.Fprintln(rw, string(output))
}

func sendError(rw http.ResponseWriter, status int) {
	rw.WriteHeader(status)
	fmt.Fprintln(rw, "Resource not found")
}