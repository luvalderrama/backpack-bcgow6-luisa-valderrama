package main

import "fmt"

func main() {
	var (
		temperatura float64
		humedad float64
		presionAtmosferica float64
	)
	temperatura, humedad, presionAtmosferica = 17.5, 123.45, 112.596
	fmt.Println(temperatura, humedad, presionAtmosferica)
}

//a las variables le asignaria un decimal pero por el momento flotante