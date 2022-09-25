package main

import "fmt"

type Data struct {
	age int
	name string
	isEmploye bool
	sueldo int
	antiguedad int
}

func main() {
	
	var client1, client2 Data
	client1.age = 23
	client1.name = "juan"
	client1.isEmploye = true
	client1.sueldo = 1000000
	client1.antiguedad = 2
	client2.age = 18
	client2.name = "jose"
	client2.isEmploye = true
	client2.sueldo = 10000000
	client2.antiguedad = 1

	slice1 := []Data{client1, client2}
	for _, slice :=  range slice1 {
    	if slice.age > 22 && slice.isEmploye == true && slice.antiguedad > 1{
			fmt.Printf("se le puede prestar a las siguientes personas: %s\n", slice.name)
			if slice.sueldo > 100000 {
				fmt.Printf("no se le cobra interes a: %s\n", slice.name)
			}
			continue
		} 
		fmt.Printf("lo sentimos no se le puede prestar a: %s\n", slice.name)


	}

}

