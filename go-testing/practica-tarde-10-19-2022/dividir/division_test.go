package dividir

import(
	"testing"
	"errors"
	"github.com/stretchr/testify/assert"

)

func TestDividir(t *testing.T) {
	a := 10
	b := 5
	esperado := 2

	resultado, err := Dividir(a, b)

	assert.Nil(t, err)
	assert.Equal(t, esperado, resultado, "el resultado es diferente del esperado")
}

func TestDividirBad(t *testing.T) {
	a := 10
	b := 0
	esperado := errors.New("el denominador no puede ser 0")

	_, err := Dividir(a, b)

	assert.NotNil(t, err)
	assert.ErrorContains(t, err, esperado.Error())
}