package main

import (
	"fmt"
	"errors"
)

func main() {
	var err = errors.New("el salario ingresado no alcanza el mínimo imponible")
	salary := 299999999
	if salary < 150000 {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Println("debes pagar impuestos")
}