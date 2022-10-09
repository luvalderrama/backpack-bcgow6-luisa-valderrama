package users

import(
	"fmt"
	"errors"
)

type User struct{
	Id int
	Nombre string
	Apellido string
	Correo string
	Edad int
	Activo bool
	FechaCreacion string
}

var usuarios  []User

type Repository interface {
	CrearUsuario(id int, nombre string, apellido string, correo string, edad int, activo bool, fechaCreacion string)(User, error)
	GetAll() ([]User)
	GetById(id int)(User, error)
	Update(id int, nombre string, apellido string, correo string, edad int, activo bool, fechaCreacion string)(User, error)
	LastId()(int)
	Delete(id int)error
	Patch(id int, nombre string, apellido string)(User, error)
}//interfaz que define los metodos que deben tener la clase que la implementen


type repository struct {
} //clase que implementa la interfaz

func NewReposirory() Repository {
	return &repository{}
}

func (r *repository) CrearUsuario(
	id int,
	nombre string,
	apellido string,
	correo string,
	edad int,
	activo bool, 
	fechaCreacion string,
) (User, error) {
	usuario := User{
		id,
		nombre,
		apellido,
		correo,
		edad,
		activo,
		fechaCreacion,
	}

	usuarios = append(usuarios, usuario)
	return usuario, nil
}

func (r *repository) GetAll() ([]User) {
	return usuarios
}

func (r *repository) GetById(id int) (User, error) {
	for _, value := range usuarios {
		if value.Id == id{
			return value, nil
		}
	}
	return User{}, errors.New("user not found")
}

func (r *repository) LastId() (int) {
	if len(usuarios) == 0 {
		return 0
	}
	return usuarios[len(usuarios)-1].Id
}

func (r *repository) Update(
	id int,
	nombre string,
	apellido string,
	correo string,
	edad int,
	activo bool,
	fechaCreacion string,
)(User, error) {
	usuario := User{
		Nombre: nombre,
		Apellido: apellido,
		Correo: correo,
		Edad: edad,
		Activo: activo,
		FechaCreacion: fechaCreacion,
	}

	update := false
	
	for i := range usuarios{
		if usuarios[i].Id == id {
			usuario.Id = id
			usuarios[i] = usuario
			update = true
			break
		}
	}

	if !update {
		return User{}, fmt.Errorf("user %d not found", id)
	}

	return usuario, nil
}

func (r *repository) Patch(id int, nombre string, apellido string) (User, error) {
	update := false
	var usuario User
	for i := range usuarios{
		if usuarios[i].Id == id {
			usuarios[i].Nombre = nombre
			usuarios[i].Apellido = apellido
			update = true
			usuario = usuarios[i]
		}
	}
	if !update {
		return User{}, fmt.Errorf("user %d not found", id)
	}
	return usuario, nil
}

func (r *repository) Delete(id int) error {
	for index := range usuarios {
		if usuarios[index].Id == id {
			usuarios = append(usuarios[:index], usuarios[index + 1:]...)
			return nil
		}
	}
	return fmt.Errorf("user %d not found", id)
}
