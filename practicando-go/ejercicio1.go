package main

import "fmt"

func getNumberType(num int) string {
	if num > 0 {
		return "the number is positive"
	}

	if num < 0 {
		return "the number is negative"
	}

	return "the number is zero"
}

func main(){
	var numberType string = getNumberType(0)

	fmt.Println(numberType)
}

