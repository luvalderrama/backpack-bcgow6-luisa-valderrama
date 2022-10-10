package main

import
(
	"log"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	"github.com/luvalderrama/backpack-bcgow6-luisa-valderrama/go-web/practica-tarde-5-10-2022/internal/users"
	"github.com/luvalderrama/backpack-bcgow6-luisa-valderrama/go-web/practica-tarde-5-10-2022/cmd/server/handler"
	"github.com/luvalderrama/backpack-bcgow6-luisa-valderrama/go-web/practica-tarde-5-10-2022/pkg/store"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error al intentar cargar archivo .env")
	}
	db := store.New(store.FileType,  "./store.json")

	repo := users.NewReposirory(db)
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