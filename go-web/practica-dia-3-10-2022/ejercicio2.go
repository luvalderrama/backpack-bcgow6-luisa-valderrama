
package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/bienvenida", func(ctx *gin.Context){
		ctx.JSON(200, gin.H{
			"message": "hola, luisa alejandra\n",
		})
	})

	router.Run()
}