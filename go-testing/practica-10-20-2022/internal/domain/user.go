package domain

type User struct{
	Id int `json:"id"`
	Nombre string `json:"nombre"`
	Apellido string `json:"apellido"`
	Correo string `json:"correo"`
	Edad int `json:"edad"`
	Activo bool `json:"activo"`
	FechaCreacion string `json:"fechaCreacion"`
}