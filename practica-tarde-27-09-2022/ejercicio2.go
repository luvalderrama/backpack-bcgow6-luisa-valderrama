package main

import "fmt"

type usuario struct {
	nombre string
	apellido string
	correo string
	productos []ProductoUsuario
}

type Producto struct {
	nombre string
	precio float32
}

type ProductoUsuario struct {
	producto Producto
	cantidad int
}

func CrearProducto(nombre string, precio float32) Producto {
	return Producto{nombre: nombre, precio: precio}
}

func (u *usuario) AgregarProducto(product ProductoUsuario) {
	u.productos = append(u.productos, product)

}

func AgregarProductoUsuario(user *usuario, producto *Producto, cantidad int) {
	productUser := ProductoUsuario{producto: *producto, cantidad: cantidad}
	user.AgregarProducto(productUser)
}

func (u *usuario) BorrarProductos() {
	u.productos = []ProductoUsuario{}
}

func BorrarProductoUsuario(user *usuario) {
	user.BorrarProductos()

}
func main() {
	producto := CrearProducto("jabon", 23344)
	usuario := usuario{
		nombre : "juan",
		apellido : "pablo",
		correo: "juan",
		productos: []ProductoUsuario{},
	}
	
	fmt.Println(usuario)
	fmt.Println(producto)

	AgregarProductoUsuario(&usuario, &producto, 3)

	fmt.Println(usuario)
	BorrarProductoUsuario(&usuario)
	fmt.Println(usuario)


}

// to do permitir agregar varios usuarios y productos, un menu