package main

import (
	"fmt"
	"log"
	"time"
	"io/ioutil"
	"net/http"
)

var req int
var apiSpec string = "/openapi.json"

func handleRequests(b string) {
	http.HandleFunc("/endpoint", func (w http.ResponseWriter, r *http.Request) {
		reqNum := req
		req += 1

		fmt.Printf("Endpoint hit number %d\n", reqNum)
		time.Sleep(3 * time.Second)
		fmt.Printf("Responding to request %d\n", reqNum)

		fmt.Fprintf(w, b)
	})

	log.Fatal(http.ListenAndServe(":10001", nil))
}

func main() {
	b, err := ioutil.ReadFile(apiSpec)
	if err != nil {
		fmt.Print(err)
		return
	}
	handleRequests(string(b))
}
