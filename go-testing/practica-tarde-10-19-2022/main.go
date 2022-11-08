package main

import(
	"fmt"
	"github.com/luvalderrama/backpack-bcgow6-luisa-valderrama/go-testing/practica-tarde-10-19-2022/calculador"
	"github.com/luvalderrama/backpack-bcgow6-luisa-valderrama/go-testing/practica-tarde-10-19-2022/ordenamiento"
	"github.com/luvalderrama/backpack-bcgow6-luisa-valderrama/go-testing/practica-tarde-10-19-2022/dividir"
)

func main() {
	a,b := 10, 5

	resta := calculador.Rest(a, b)
	fmt.Println(resta)

	slice := []int{5, 3, 4, 7, 8, 9}
	ordenar := ordenamiento.Order(slice)
	fmt.Println(ordenar)

	division, _ := dividir.Dividir(a, b)
	fmt.Println(division)
}