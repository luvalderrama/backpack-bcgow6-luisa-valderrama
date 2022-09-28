package main

import (
	"fmt"
)

func main() {
	salary := 299
	if salary < 150000 {
		err := fmt.Errorf("el salario ingresado no alcanza el mínimo imponible %d", salary)
		fmt.Println(err)
		return
	}
	fmt.Println("debes pagar impuestos")
}