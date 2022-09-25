package main

import "fmt"

func main() {
	var employees = map[string]int{
	"Benjamin": 20,
	"Nahuel": 26,
	"Brenda": 19,
	"Darío": 44, 
	"Pedro": 30}
	var counter int
	for value := range employees { 
		if employees[value] > 21{
			counter ++
		}
	}
	fmt.Println("total de empleados mayores a 21 son: ", counter)
	employees["federico"] = 25
	delete(employees, "Pedro")
	fmt.Println(employees)
}