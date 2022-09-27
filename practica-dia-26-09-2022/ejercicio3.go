package main

import "fmt"

type CategoryFunction func(float32) float32
type CategoriesMap map[string]CategoryFunction

var categories = CategoriesMap{
	"A": func(workedMinutes float32) float32 { return calculateSalary(3000, workedMinutes, 50) },
	"B": func(workedMinutes float32) float32 { return calculateSalary(1500, workedMinutes, 20) },
	"C": func(workedMinutes float32) float32 { return calculateSalary(1000, workedMinutes, 0) },
}

func calculateSalary(moneyPerHour, workedMinutes, increment float32) float32 {
	var workedHours float32 = workedMinutes / 60

	return workedHours * moneyPerHour * (increment / 100 + 1)
}

func getEmployeeSalary(category string, workedMinutes float32) float32 {
	var salary float32 = categories[category](workedMinutes)
	return salary
}

func main() {
	var employeeSalary float32 = getEmployeeSalary("B", 1500)

	fmt.Println(employeeSalary)
}

// TO DO: SOPORTAR ERRORES