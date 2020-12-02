package main

var lstCity []city

//main entry point
func main() {
	lstCity = loadCityListFile()
	handleRequests()
}
