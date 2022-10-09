
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)



type user struct{
	Id string `json:"id"`
	Nombre string `json:"nombre"`
	Apellido string `json:"apellido"`
	Correo string `json:"correo"`
	Edad int `json:"edad"`
	Activo string `json:"activo"`
	FechaCreacion string `json:"fecha"`
}

var users = []user {
	{
		Id: "1",
		Nombre: "luisa",
		Apellido: "valderrama",
		Correo : "luisa@gmail.com",
		Edad : 19,
		Activo: "activo",
		FechaCreacion : "03-10-2022",
	},
	{
		Id: "2",
		Nombre: "luis",
		Apellido: "valderrama",
		Correo : "luis@gmail.com",
		Edad : 30,
		Activo: "inactive",
		FechaCreacion : "03-10-2022",
	},
}

func GetAll(ctx *gin.Context) {
    ctx.IndentedJSON(http.StatusOK, users)
}

func main() {
	router := gin.Default()

	router.GET("/usuarios", GetAll)
	router.Run()
}
