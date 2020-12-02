package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

//http response
func searchByCity(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["city"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'city' is missing")
		return
	}

	results := doSearch(keys[0])
	resultText := ""
	resultNames := make([]string, len(results))

	for i, element := range results {
		resultText += strings.Join([]string{strconv.Itoa(i), ": ", element.Name, "\n"}, "")
		resultNames[i] = element.Name
	}

	//send json back
	returnVal, _ := json.Marshal(resultNames)

	w.Header().Set("Content-Type", "application/json")
	w.Write(returnVal)
	fmt.Println("Endpoint Hit: homePage; returned: " + resultText)
}

//setup routes
func handleRequests() {
	http.HandleFunc("/", searchByCity)
	log.Fatal(http.ListenAndServe(":10000", nil))
}
