package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

var lstCity []city

func homePage(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["key"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}

	results := parseCityJSON(keys[0])
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

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	lstCity = loadCityListFile()
	handleRequests()
}

//json stuff
type city struct {
	ID      int
	Name    string
	Country string
	Coord   coordinates
}
type coordinates struct {
	Lon, Lat float64
}

func loadCityListFile() []city {
	log.Println("Started loading file!")

	ret := readCityListFile()

	log.Println("Started sorting file!")

	//sort by city name, ties are left in original order
	sort.SliceStable(ret, func(i, j int) bool {
		//return ret[i].Name < ret[j].Name
		return strings.ToLower(ret[i].Name) < strings.ToLower(ret[j].Name)
	})

	log.Println("Finished loading file!")

	return ret
}

func readCityListFile() []city {
	file, _ := ioutil.ReadFile("resources/city.list.json")

	ret := make([]city, 0)

	_ = json.Unmarshal([]byte(file), &ret)

	return ret
}

func parseCityJSON(searchVal string) []city {
	ret := make([]city, 0)

	idx := sort.Search(len(lstCity), func(idx int) bool { return searchStringsStartsWith(searchVal, idx) })
	fmt.Println(idx)

	//if no match, search returns ending index of search range
	if idx != len(lstCity) {
		fmt.Println(lstCity[idx].Name)
		ret = append(ret, lstCity[idx])
	}

	//traverse the next 4 results and add them if needed
	for n := 0; n < 400; n++ {
		idx++
		if checkNextResult(searchVal, idx, lstCity) {
			ret = append(ret, lstCity[idx])
		} else {
			return ret
		}
	}
	return ret
}

func searchStringsStartsWith(val string, idx int) bool {
	sliceLen := 0
	if len(lstCity[idx].Name) < len(val) {
		sliceLen = len(lstCity[idx].Name)
	} else {
		sliceLen = len(val)
	}

	//ToLower both sides for a case insensitive compare
	//return val <= lstCity[idx].Name[:sliceLen]
	return strings.ToLower(val) <= strings.ToLower(lstCity[idx].Name[:sliceLen])
}

func checkNextResult(val string, idx int, lstCity []city) bool {
	//return searchStringsStartsWith(searchVal, idx)
	sliceLen := 0
	if len(lstCity[idx].Name) < len(val) {
		sliceLen = len(lstCity[idx].Name)
	} else {
		sliceLen = len(val)
	}

	return strings.ToLower(val) == strings.ToLower(lstCity[idx].Name[:sliceLen])
}
