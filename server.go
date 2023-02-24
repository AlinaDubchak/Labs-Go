package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/time", showTime)
	err := http.ListenAndServe(":8895", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func showTime(write http.ResponseWriter, request *http.Request) {
	timeNowFormatRFC3339 := time.Now().Format(time.RFC3339)
	timeData := map[string]string{"time": timeNowFormatRFC3339}
	respond, err := json.MarshalIndent(timeData, "", "")
	if err != nil {
		log.Fatal(err)
	}
	if request.Method == "GET" {
		fmt.Fprintf(write, string(respond))
	}
}
