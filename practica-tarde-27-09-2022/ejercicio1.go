package main

import "fmt"

type Usuario struct {
	nombre string
	apellido string
	edad int
	correo string
	contrasenia string
}

func (u *Usuario) CambiarNombre(nombre, apellido string) {
	u.nombre = nombre
	u.apellido = apellido
}

func (u *Usuario) CambiarEdad(edad int) {
	u.edad = edad
}

func (u *Usuario) CambiarCorreo(correo string) {
	u.correo = correo
}

func (u *Usuario) CambiarContrasenia(contrasenia string) {
	u.contrasenia = contrasenia
}

func main() {
	usuario := Usuario{
		nombre: "luisa",
		apellido: "valderrama",
		edad: 18,
		correo: "luvalde@gmail.com",
		contrasenia: "jsjdkk",
	}
	fmt.Println(usuario)
	usuario.CambiarNombre("juan", "pablo")
	fmt.Println(usuario)
}