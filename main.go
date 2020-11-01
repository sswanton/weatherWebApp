// Starter code taken from https://tutorialedge.net/golang/creating-restful-api-with-golang/

package main

import (
	"fmt"
	"log"
	"net/http"
)

func mainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
	fmt.Println("Endpoint Hit: mainPage")
}

func handleRequests() {
	http.HandleFunc("/", mainPage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
