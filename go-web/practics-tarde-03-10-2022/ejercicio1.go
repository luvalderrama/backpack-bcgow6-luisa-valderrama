
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"fmt"
)



type User struct{
	Id int `json:"id"`
	Nombre string `json:"nombre"`
	Apellido string `json:"apellido"`
	Correo string `json:"correo"`
	Edad int `json:"edad"`
	Activo bool `json:"activo"`
	FechaCreacion string `json:"fechaCreacion"`
}

var users = []User {
	{
		Id: 1,
		Nombre: "luisa",
		Apellido: "valderrama",
		Correo: "luisa@gmail.com",
		Edad: 19,
		Activo: true,
		FechaCreacion: "03-10-2022",
	},
	{
		Id: 2,
		Nombre: "luis",
		Apellido: "valderrama",
		Correo: "luis@gmail.com",
		Edad: 30,
		Activo: false, 
		FechaCreacion: "03-10-2022",
	},
}

func BuscarUsuarios(ctx *gin.Context) {
    nombre := ctx.Query("nombre")
		usuariosFiltrados := []User{}
    for _, value := range users {
				if nombre != "" && strings.Contains(value.Nombre, nombre){
					usuariosFiltrados = append(usuariosFiltrados, value)
				}
    }
		fmt.Println(usuariosFiltrados)

    ctx.JSON(http.StatusOK, &usuariosFiltrados)
}


func main() {
	router := gin.Default()
	router.GET("/usuarios/filter", BuscarUsuarios)
	router.Run()
}