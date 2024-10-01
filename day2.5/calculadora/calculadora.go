package calculadora

import (
	"fmt"
)


func soma(a, b float64) float64 {
	return a + b
}

func subtrai(a, b float64) float64 {
	return a - b
}

func multiplique(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) float64 {
	if b == 0 {
		fmt.Println("Erro: Divisão por zero!")
		return 0
	}
	return float64(a) / float64(b)
}


func Calculadora (a, b float64, operador string) (float64, error)  {
	switch operador {
	case "+":
		return soma (a, b), nil
	case "-":
		return subtrai (a, b), nil
	case "*":
		return multiplique (a, b), nil
	case "/":
		resultado := divide(a,b)
		if resultado == 0 {
			return 0, fmt.Errorf("divisão por zero")
		}
		return resultado, nil
	default:
		return 0, fmt.Errorf("operador inválido: %s", operador)
	}
}