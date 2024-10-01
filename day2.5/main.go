package main

import (
	"fmt"
	"strconv"
	"calculadora/calculadora"
)

func obterNumero (prompt string) float64 {
	var input string
	for {
		fmt.Print(prompt)
		fmt.Scan(&input)
		num, err := strconv.ParseFloat(input, 64)
		if err == nil  {
			return num;
		}
		fmt.Println("Entrava inválida, digite um número válido")
	}
}

func obterOperador () string {
	var operador string
	for {
	fmt.Print("Digite um operador entre as opções(+  -  /  *): ")
	fmt.Scan(&operador)
		switch operador {
		case "+":
			return "+"
		case "-":
			return "-"
		case "*":
			return "*"
		case "/": 
			return "/"
		default:
			fmt.Println("Entrada inválida, por favor escolha uma das operações (+  -  /  *) ")
		}
	}
}

func main () {
	num1 := obterNumero("Digite o primeiro número: ")
	num2 := obterNumero("Digite o segundo número: ")
	operador := obterOperador()

	resultado, err := calculadora.Calculadora(num1, num2, operador) 
	if err != nil{
		fmt.Printf("Erro ao calcular a operação %.2f %s %.2f \n", num1, operador, num2)
		return
	}

	if resultado == float64(int(resultado)) {
		fmt.Printf("Resultado: %d\n", int(resultado))
	} else {
		fmt.Printf("Resultado: %.2f\n", resultado)
	}
}