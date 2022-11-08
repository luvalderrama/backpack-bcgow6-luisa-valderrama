package main

import(
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	"github.com/luvalderrama/backpack-bcgow6-luisa-valderrama/go-testing/practica-10-20-2022/internal/users"
	"github.com/luvalderrama/backpack-bcgow6-luisa-valderrama/go-testing/practica-10-20-2022/docs"
	"github.com/luvalderrama/backpack-bcgow6-luisa-valderrama/go-testing/practica-10-20-2022/cmd/server/handler"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)


// @title           Practica 6
// @version         1.0
// @description     users crud.
// @termsOfService  https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones
// @contact.name   API Support Digital House
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host      localhost:8080
// @BasePath  /api/v1
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error al intentar cargar archivo .env")
	}

	repo := users.NewReposirory()
	service := users.NewService(repo)
	userHandler := handler.NewUser(service)

	router := gin.Default()

	api := router.Group("/api/v1")
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	api.GET("./docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	userGroup := api.Group("/usuarios")
	userGroup.POST("/", userHandler.CrearUsuario())
	userGroup.GET("/all", userHandler.GetAll())
	userGroup.PUT("/:id", userHandler.Update())
	userGroup.PATCH("/:id", userHandler.Patch())
	userGroup.GET("/:id", userHandler.GetById())
	userGroup.DELETE("/:id", userHandler.Delete())

	router.Run()
}