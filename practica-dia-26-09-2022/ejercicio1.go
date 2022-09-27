package main

import (
	"errors"
	"fmt"
)

const (
	MinSalario = 0.17
	MaxSalario = 0.27
)

func calcularImpuestos(sueldo float32)(float32, error){
	if sueldo < 0 {
		return sueldo, errors.New("no se puede ingresar un numero menor a 0")
	}

	if sueldo > 50000 && sueldo <= 150000 {
		var retencion float32 = sueldo * MinSalario
		
		return retencion, nil
	}

	if sueldo > 150000 {
		var retencion float32 = sueldo * MaxSalario

		return retencion, nil
	}

	return 0, nil
}

func main() {
	impuestos, err  := calcularImpuestos(0)

	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(impuestos)
}