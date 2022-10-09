
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type user struct{
	Id int `json:"id"`
	Nombre string `json:"nombre"`
	Apellido string `json:"apellido"`
	Correo string `json:"correo"`
	Edad int `json:"edad"`
	Activo bool `json:"activo"`
	FechaCreacion string `json:"fechaCreacion"`
}

var usuarios = []user {
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

func GetId(ctx *gin.Context) {
	id := ctx.Param("id")
	idNumber, err := strconv.Atoi(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "el id no es valido"})
		return
	}

	for _, value := range usuarios {
		if value.Id == idNumber {
			ctx.JSON(http.StatusOK, &value)
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"msg": "usuario no encontrado"})
}

func main() {
	router := gin.Default()
	router.GET("/usuarios/:id", GetId)

	router.Run()
}
