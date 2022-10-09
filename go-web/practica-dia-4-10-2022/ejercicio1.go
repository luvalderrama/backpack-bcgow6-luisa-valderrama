package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

type User struct{
	Id int `json:"id"`
	Nombre string `json:"nombre" binding:"required"`
	Apellido string `json:"apellido" binding:"required"`
	Correo string `json:"correo" binding:"required"`
	Edad int `json:"edad" binding:"required"`
	Activo bool `json:"activo" binding:"required"`
	FechaCreacion string `json:"fechaCreacion" binding:"required"`
}

var usuarios = []User{}

func CrearUsuario() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "luisuiie23" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "no tiene permiso para realizar la peticion solicitada"})
		}

		var user User
		err := ctx.ShouldBindJSON(&user)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user.Id = len(usuarios) + 1
		usuarios = append(usuarios, user)
		fmt.Println(usuarios)
		ctx.JSON(http.StatusOK, user)
	}
}

func main() {
	router := gin.Default()
	newRouter := router.Group("/usuarios")
	newRouter.POST("/", CrearUsuario())
	router.Run()
}