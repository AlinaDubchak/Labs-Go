package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

const PORT, PATH string = ":8795", "/time"

func main() {
	http.HandleFunc(PATH, showTime)
	err := http.ListenAndServe(PORT, nil)
	http.HandleFunc("/time", showTime)
	log.Fatal(http.ListenAndServe(":8795", nil))
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
