package main

import (
	"errors"
	"fmt"
)

func calcularPromedio(notas ...float32) (float32, error) {
	var totalNotas float32 = 0
	var tieneError bool = false

	for _, nota := range notas {
		if nota < 0 {
			tieneError = true
			break
		}

		totalNotas += nota
	}

	if tieneError {
		return 0, errors.New("no se puede ingresar una nota menor a 0")
	}

	var promedio float32 = totalNotas / float32(len(notas))

	return promedio, nil
}

func main() {
	promedio, err := calcularPromedio(4.5, 3.0, 5.0)

	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(promedio)

}