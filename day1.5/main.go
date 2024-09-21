package main

import "fmt"

// const prefixoOlaPortugues = "Olá, "

func Ola(nome string, lingua string) string{
	prefixoOla := "veja so"
	if nome == "" {
		nome = "mundo!"
	}
	if lingua != "ingles" && lingua != "espanhol" || lingua == "" {
		prefixoOla = "Olá, "
	}
	if lingua == "espanhol" {
		prefixoOla = "Hola, "
	}
	if lingua == "ingles" {
		prefixoOla = "Hi, "
	}
	return prefixoOla + nome
}

func main()  {
	fmt.Println(Ola("Daniel", "espanhol"))
}
