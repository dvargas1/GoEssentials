package main

import "testing"

func TestOla(t *testing.T) {
	verificarResultado := func (t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado %q, esperado %q", resultado, esperado)
		}
	}
	t.Run("Diga Olá para: ", func (t *testing.T) {
		resultado := Ola("Daniel", "")
		esperado := "Olá, Daniel"
		verificarResultado(t, resultado, esperado)
	})
	t.Run("Diga Olá em espanhol para: ", func (t *testing.T) {
		resultado := Ola("Daniel", "espanhol")
		esperado := "Hola, Daniel"
		verificarResultado(t, resultado, esperado)
	})
	t.Run("Diga Olá em ingles para: ", func (t *testing.T) {
		resultado := Ola("Daniel", "ingles")
		esperado := "Hi, Daniel"
		verificarResultado(t, resultado, esperado)
	})
	t.Run("Se a string estiver vazia, chamamos um Ola mundo", func(t *testing.T) {
		resultado := Ola("", "")
		esperado := "Olá, mundo!"
		verificarResultado(t, resultado, esperado)
	})
}
