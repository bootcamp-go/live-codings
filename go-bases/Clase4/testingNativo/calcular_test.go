package calcular

// Se importa el package testing
import (
	"testing"
)

func TestAddition(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	num1 := 3
	num2 := 5
	expecteResult := 8

	// Se ejecuta el test
	result := Addition(num1, num2)

	// Se validan los resultados
	if result != expecteResult {
		t.Errorf("Funcion Addition() arrojo el resultado = %v, pero el esperado es  %v", result, expecteResult)
	}

}

func TestSustraction(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)
	num1 := 3
	num2 := 5
	expecteResult := -2

	// Se ejecuta el test
	result := Sustraction(num1, num2)

	// Se validan los resultados
	if result != expecteResult {
		t.Errorf("Funcion Sustraction() arrojo el resultado = %v, pero el esperado es  %v", result, expecteResult)
	}
}
