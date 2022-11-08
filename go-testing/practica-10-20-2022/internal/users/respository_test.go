package users

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
)

type MockUser struct {
	dataMock []User
	
}

func (m *MockUser) read(data interface{}) (err error) {
	casteData := data.(*[]User)
	*casteData = m.dataMock
	return nil
}

func (m *MockUser) write(data interface{})(err error) {
	casteData := data.([]User)
	m.dataMock = append(m.dataMock, casteData[len(casteData)-1])
	return nil
}

func TestGetAll(t *testing.T) {
	dataBase := []User{
		{
			Id: 1,
			Nombre: "luisa",
			Apellido: "marin",
			Correo: "luidsbb@gmail.com",
			Edad: 21,
			Activo: true,
			FechaCreacion: "12/29/2000",
		},
	}
	repository := NewReposirory()
	_, err := repository.CrearUsuario(1, "luisa", "marin", "luidsbb@gmail.com", 21, true, "12/29/2000")
	if err != nil {
		fmt.Println(err)
		return
	}


	mockUser := MockUser{
		dataMock: dataBase,
	}


	mockUser.dataMock = append(mockUser.dataMock, User{})


	result := repository.GetAll()
	assert.Nil(t, err)
	assert.Equal(t, dataBase, result)
}