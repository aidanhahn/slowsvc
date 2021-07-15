package main

import (
	"fmt"
	"log"
	"time"
	"net/http"
)

var req int
var apiSpec string = `
{
  "openapi": "3.0.0",
  "info": {
    "title": "API Spec API",
    "version": "1.0.0"
  },
  "paths": {
    "/": {
      "get": {
        "operationId": "get this spec",
        "summary": "Returns this specification",
        "responses": {
          "200": {
            "description": "200 response",
            "content": {
              "schema": {
                type: object
              }
            }
          }
        }
      }
    }
  }
}
`

func endpoint(w http.ResponseWriter, r *http.Request){
	reqNum := req
	req += 1

	fmt.Printf("Endpoint hit number %d\n", reqNum)
	time.Sleep(3 * time.Second)
	fmt.Printf("Responding to request %d\n", reqNum)

	fmt.Fprintf(w, apiSpec)
}

func handleRequests() {
	http.HandleFunc("/endpoint", endpoint)
	log.Fatal(http.ListenAndServe(":10001", nil))
}

func main() {
	handleRequests()
}
