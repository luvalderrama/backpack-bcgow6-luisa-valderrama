package handler

import(
	"os"
	"strconv"
	"net/http"
	"github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gin-gonic/gin"
	"github.com/luvalderrama/backpack-bcgow6-luisa-valderrama/go-web/practica-tarde-6-10-2022/internal/users"
	"github.com/luvalderrama/backpack-bcgow6-luisa-valderrama/go-web/practica-tarde-6-10-2022/pkg/web"
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

func (r *request) ValidatePatch() error {
	return validation.ValidateStruct(r,
		validation.Field(&r.Nombre, validation.Required),
		validation.Field(&r.Apellido, validation.Required),
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

//CrearUsuario User godoc
// @Summary  CrearUsuario User
// @Tags     users
// @Produce  json
// @Param    token    header    string          true  "token invalido"
// @Param    user     body      request  true  "User to CrearUsuario"
// @Success  200      {object}  web.Response
// @Failure  401      {object}  web.Response
// @Failure  400      {object}  web.Response
// @Failure  500      {object}  web.Response
// @Router   /usuarios [POST]
func (c *User) CrearUsuario() gin.HandlerFunc {
	return func(ctx *gin.Context){
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "token invalido"))
		}
		var req request
		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}
		err = req.Validate()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}

		usuario, err := c.service.CrearUsuario(req.Nombre, req.Apellido, req.Correo, req.Edad, req.Activo, req.FechaCreacion)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, usuario, ""))
	}

}

// Update User godoc
// @Summary  Update User
// @Tags     users
// @Produce  json
// @Param    token    header    string          true  "token invalido"
// @Param    id       path      int             true   "Id User"
// @Param    user     body      request  true  "User to CrearUsuario"
// @Success  200      {object}  web.Response
// @Failure  401      {object}  web.Response
// @Failure  400      {object}  web.Response
// @Failure  500      {object}  web.Response
// @Router   /usuarios/{id} [PUT]
func (c *User) Update() gin.HandlerFunc {
	return func(ctx *gin.Context){
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "token invalido"))
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}
		var req request
		err = ctx.ShouldBindJSON(&req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}
		err = req.Validate()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}
		usuario, err := c.service.Update(int(id), req.Nombre, req.Apellido, req.Correo, req.Edad, req.Activo, req.FechaCreacion)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, usuario, ""))
	}
}

// patch User godoc
// @Summary  patch User
// @Tags     users
// @Produce  json
// @Param    token    header    string          true  "token invalido"
// @Param    id       path      int             true   "Id User"
// @Param    user     body      request  true  "User to Patch"
// @Success  200      {object}  web.Response
// @Failure  401      {object}  web.Response
// @Failure  400      {object}  web.Response
// @Failure  500      {object}  web.Response
// @Router       /usuarios/{id} [PATCH]
func (c *User) Patch() gin.HandlerFunc {
	return func(ctx *gin.Context){
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "token invalido"))
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}
		var req request
		err = ctx.ShouldBindJSON(&req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}

		err = req.ValidatePatch()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}
		usuario, err := c.service.Patch(int(id), req.Nombre, req.Apellido)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, usuario, ""))
	}
}

// ListUsuarios godoc
// @Summary  GetAll User
// @Tags     users
// @Produce  json
// @Param    token  header    string        true  "token"
// @Success  200    {object}  web.Response  "List usuarios"
// @Failure  401    {object}  web.Response  "Unauthorized"
// @Router   /usuarios/all [GET]
func (c *User) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "token invalido"))
			return
		}
		usuarios := c.service.GetAll()
		ctx.JSON(http.StatusOK, usuarios)
	}
}

// GetById User
// @Summary  GetById User
// @Tags     users
// @Param    id     path      int     true  "Id User"
// @Param    token  header    string  true  "Token"
// @Success  200    {object}  web.Response
// @Failure  401    {object}  web.Response
// @Failure  400    {object}  web.Response
// @Failure  404    {object}  web.Response
// @Router   /usuarios/{id} [GET]
func (c *User) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "token invalido"))
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}
		usuario, err := c.service.GetById(int(id))
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, usuario, ""))
	}
}

// Delete User
// @Summary  Delete User
// @Tags     users
// @Param    id     path      int     true  "Id User"
// @Param    token  header    string  true  "Token"
// @Success  200    {object}  web.Response
// @Failure  401    {object}  web.Response
// @Failure  400    {object}  web.Response
// @Failure  404    {object}  web.Response
// @Router   /usuarios/{id} [DELETE]
func (c *User) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "token invalido"))
			return
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}
		err = c.service.Delete((int(id)))
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
		}
		
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, nil, "usuario fue eliminado"))
	}
}