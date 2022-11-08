package calculador

import(
	"testing"
	"github.com/stretchr/testify/assert"

)

func TestSub(t *testing.T) {
	a := 10
	b := 5
	esperado := 5

	resultado := 	Rest(a, b)

	assert.Equal(t, esperado, resultado, "el resultado es diferente del esperado")
}