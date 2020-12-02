package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

//main json parser entry point
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

//subroutine for reading file and parsing into city objects
func readCityListFile() []city {
	return deserializeFile(readFile())
}

func readFile() []byte {
	file, _ := ioutil.ReadFile("resources/city.list.json")

	return file
}

func deserializeFile(file []byte) []city {
	ret := make([]city, 0)

	_ = json.Unmarshal([]byte(file), &ret)

	return ret
}
