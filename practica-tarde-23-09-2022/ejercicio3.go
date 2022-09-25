package main

import (
	"fmt"
)

func main() {
	number := 10
	slice1 := map[string]interface{}{"enero": 1,
	"febrero": 2,
	"marzo": 3,
	"abril": 4,
	"mayo": 5,
	"junio": 6,
	"julio": 7,
	"agosto": 8,
	"septiembre": 9,
	"octubre": 10,
	"noviembre": 11,
	"diciembre": 12}


	for key, value := range slice1 {
		if value == number{
			fmt.Printf("%s\n", key)
		}
	}
}

// se puede resolver de otra manera, por ejemplo usando un for y con el llenando los datos del dict que se autollene si escrbir uno por uno el mes