package dividir

import(
	"errors"
)

func Dividir(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("el denominador no puede ser 0")
	}

	return a / b , nil
}