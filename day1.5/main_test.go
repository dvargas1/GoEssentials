package main

import "testing"

func TestOla(t *testing.T) {
	resultado := Ola("Daniel")
	esperado := "Olá, Daniel"

	if resultado != esperado {
		t.Errorf("resultado %q, esperado %q", resultado, esperado)
	}
}
