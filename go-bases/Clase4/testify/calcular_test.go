package calcular

// Se importa el package testing
import (
	"testing"

	"github.com/stretchr/testify/assert" // Se importa testify
)

func TestAddition(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	num1 := 3
	num2 := 5
	expecteResult := 8

	// Se ejecuta el test
	result := Addition(num1, num2)

	// Se validan los resultados
	assert.Equal(t, result, expecteResult)
}

func TestSustraction(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	num1 := 3
	num2 := 5
	expecteResult := -2

	// Se ejecuta el test
	result := Sustraction(num1, num2)

	// Se validan los resultados
	assert.Equal(t, result, expecteResult)
}
