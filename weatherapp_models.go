package main

//json structs
type city struct {
	ID      int
	Name    string
	Country string
	Coord   coordinates
}
type coordinates struct {
	Lon, Lat float64
}
