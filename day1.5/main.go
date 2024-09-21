package main

import "fmt"

// const prefixoOlaPortugues = "Olá, "

func Ola(nome string, lingua string) string{
	prefixoOla := "Olá, "
	if nome == "" {
		nome = "mundo!"
	}
	switch lingua{
		case "espanhol":
			prefixoOla = "Hola, "
		case "ingles":
			prefixoOla = "Hi, "
		}
	return prefixoOla + nome
}

func main()  {
	fmt.Println(Ola("Daniel", "espanhol"))
}
