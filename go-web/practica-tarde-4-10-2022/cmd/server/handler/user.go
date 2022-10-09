package handler

import(
	"strconv"
	"net/http"
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gin-gonic/gin"
	"github.com/luvalderrama/backpack-bcgow6-luisa-valderrama/go-web/practica-tarde-4-10-2022/internal/users"
)

type request struct{
	Id int `json:"id"`
	Nombre string `json:"nombre"`
	Apellido string `json:"apellido"`
	Correo string `json:"correo"`
	Edad int `json:"edad"`
	Activo bool `json:"activo"`
	FechaCreacion string `json:"fechaCreacion"`
}

func (r *request) Validate() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Nombre, validation.Required),
		validation.Field(&r.Apellido, validation.Required),
		validation.Field(&r.Correo, validation.Required),
		validation.Field(&r.Edad, validation.Required),
		validation.Field(&r.Activo, validation.Required),
		validation.Field(&r.FechaCreacion, validation.Required),
	)
}

type User struct {
	service users.Service
	// nombre que le llamo al package que se llama users y Service es el contracto
}

func NewUser(u users.Service) *User {
	return &User{
		service: u,
	}
}

func (c *User) CrearUsuario() gin.HandlerFunc {
	return func(ctx *gin.Context){
		token := ctx.GetHeader("token")
		if token != "luisuiie23" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "no tiene permiso para realizar la peticion solicitada"})
			return
		}
		var req request
		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err = req.Validate()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"mesage": err.Error()})
			return
		}

		usuario, err := c.service.CrearUsuario(req.Nombre, req.Apellido, req.Correo, req.Edad, req.Activo, req.FechaCreacion)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, usuario)
	}

}

func (c *User) Update() gin.HandlerFunc {
	return func(ctx *gin.Context){
		token := ctx.GetHeader("token")
		if token != "luisuiie23" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "no tiene permiso para realizar la peticion solicitada"})
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
			return
		}
		var req request
		err = ctx.ShouldBindJSON(&req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err = req.Validate()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"mesage": err.Error()})
			return
		}
		usuario, err := c.service.Update(int(id), req.Nombre, req.Apellido, req.Correo, req.Edad, req.Activo, req.FechaCreacion)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, usuario)
	}
}

func (c *User) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "luisuiie23" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "no tiene permiso para realizar la peticion solicitada"})
			return
		}
		usuarios := c.service.GetAll()
		ctx.JSON(http.StatusOK, usuarios)
	}
}

func (c *User) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "luisuiie23" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "no tiene permiso para realizar la peticion solicitada"})
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
			return
		}
		usuario, err := c.service.GetById(int(id))
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, usuario)
	}
}