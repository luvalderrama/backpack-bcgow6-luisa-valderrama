package main

import "fmt"

func main() {
  var (
	nombre string
	apellido string
	edad int
	estaturaDeLaPersona float64
  )	 

  nombre = "luisa"
  apellido  = "valderrama"
  edad = 21
  apellidoSegundo := "marin"
  licenciaDeConducir := true
  estaturaDeLaPersona = 1.68
  cantidadDeHijos := 2

  fmt.Println(nombre, apellido, edad, apellidoSegundo, licenciaDeConducir, estaturaDeLaPersona, cantidadDeHijos)

}

//el nombre de una variable 1nombre no se puede iniciar por un numero.

// edad tambien esta mal puesto que primero es la variable.

// licencia_de_conducir no es la mejor forma puesto que esta usando _ tanbien si va inicializar una variable con valor debe usar :=

// estatura esta mal puesto que la variable no puede tener espacios
