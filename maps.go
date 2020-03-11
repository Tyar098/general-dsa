package main

import "fmt"

type Point struct {
	Lat, Long float64
}

var m map[string]Point

func main() {
	m = make(map[string]Point)
	m["Bandung"] = Point{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bandung"])
}
