package main

import "fmt"

var kelvinEbulitionPoint float64 = 373.2

func main() {
	celsiusEbulitionPoint := kelvinEbulitionPoint - 273.0

	fmt.Println("Ponto de ebulição da água em C: ", celsiusEbulitionPoint)
}
