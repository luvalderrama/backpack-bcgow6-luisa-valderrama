package main

import (
	"fmt"
)

type MyCustonErro struct {
	mesage string
}

func main() {
	salary := 299999999
	err := MyCustonErro{mesage: " el salario ingresado no alcanza el mínimo imponible"}
	if salary < 150000 {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Println("debes pagar impuestos")
}