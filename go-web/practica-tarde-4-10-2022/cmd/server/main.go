package main

import(
	"github.com/gin-gonic/gin"
	"github.com/luvalderrama/backpack-bcgow6-luisa-valderrama/go-web/practica-tarde-4-10-2022/internal/users"
	"github.com/luvalderrama/backpack-bcgow6-luisa-valderrama/go-web/practica-tarde-4-10-2022/cmd/server/handler"
)

func main() {
	repo := users.NewReposirory()
	service := users.NewService(repo)
	userHandler := handler.NewUser(service)

	router := gin.Default()

	userGroup := router.Group("/usuarios")
	userGroup.POST("/", userHandler.CrearUsuario())
	userGroup.GET("/all", userHandler.GetAll())
	userGroup.PUT("/:id", userHandler.Update())
	userGroup.PATCH("/:id", userHandler.Patch())
	userGroup.GET("/:id", userHandler.GetById())
	userGroup.DELETE("/:id", userHandler.Delete())

	router.Run()
}