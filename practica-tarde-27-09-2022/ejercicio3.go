package main

import "fmt"

type Product struct {
	nombre string
	precio float32
	cantidad int
}

type Servicios struct {
	nombre string
	precio float32
	minutosTrabajado  float32
}

type Mantenimiento struct {
	nombre string
	precio float32
}

func CalcularProductos(producto *[]Product, canal chan float32)float32 {
	var total float32
	for _, value := range *producto {
		total += value.precio * float32(value.cantidad)
	}
	canal <- total //sobreescribiendo
	return total
}

func CalcularServicios(servicios *[]Servicios, canal chan float32)float32 {
	var total float32
	for _, value := range *servicios {
		if value.minutosTrabajado < 30 {
			total += value.minutosTrabajado * 30
			fmt.Println(total)
			continue
		}
		total += value.minutosTrabajado * value.precio
	}
	canal <- total
	return total
}

func CalcularMantenimiento(matenimiento *[]Mantenimiento, canal chan float32)float32{
	var total float32
	for _, value := range *matenimiento {
		total += value.precio
	}
	canal <- total
	return total
}
func main() {
	producto := []Product{
		{nombre: "juan", precio: 255555, cantidad: 2},
		{nombre: "luisa", precio: 700000, cantidad: 3},

	}

	servicios := []Servicios{
		{nombre: "juan", precio: 255555, minutosTrabajado: 30},
		{nombre: "luisa", precio: 700000, minutosTrabajado: 40},
	}

	mantemiento := []Mantenimiento{
		{nombre: "juan", precio: 255555},
		{nombre: "luisa", precio: 700000},
	}

	canal1 := make(chan float32)
	canal2 := make(chan float32)
	canal3 := make(chan float32)

	go CalcularProductos(&producto, canal1)
	go CalcularServicios(&servicios, canal2)
	go CalcularMantenimiento(&mantemiento, canal3)

	LerrCanal1 := <-canal1 //leyendo
	LerrCanal2 := <-canal2
	LerrCanal3 := <-canal3
	total := LerrCanal1 + LerrCanal2 + LerrCanal3
	fmt.Println("total: ", total)
}