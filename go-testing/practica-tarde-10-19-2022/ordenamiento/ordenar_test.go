package ordenamiento

import(
	"testing"
	"github.com/stretchr/testify/assert"

)

func TestOrder(t *testing.T) {
	slice := []int{5, 3, 4, 7, 8, 9}
	esperado := []int{3, 4, 5, 7, 8, 9}

	resultado := 	Order(slice)

	assert.Equal(t, esperado, resultado, "el resultado es diferente del esperado")
}