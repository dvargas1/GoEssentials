package main

import "fmt"

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
	return a / b
}

func main()  {
	fmt.Println(soma(1,1))
	fmt.Println(multiplique(5,20))
	fmt.Println(divide(4,2))
	fmt.Println(subtrai(10,10))
}