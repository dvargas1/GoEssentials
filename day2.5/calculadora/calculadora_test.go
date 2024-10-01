package calculadora

import (
    "testing"
)

// Testa a função Calculadora da calculadora
func TestCalculadora(t *testing.T) {
    // Função auxiliar para verificar o resultado de cada teste
    verificarResultado := func(t *testing.T, resultado, esperado float64, err error, deveErrar bool) {
        t.Helper() // essa função é uma ajudante de teste
        if (err != nil) != deveErrar {
            t.Errorf("Erro: %v, deveErrar: %v", err, deveErrar)
        }
        if resultado != esperado && !deveErrar {
            t.Errorf("resultado %v, esperado %v", resultado, esperado)
        }
    }

    // Teste de Soma
    t.Run("Soma de 5 e 2", func(t *testing.T) {
        resultado, err := Calculadora(5, 2, "+")
        verificarResultado(t, resultado, 7, err, false)
    })

    // Teste de Subtração
    t.Run("Subtração de 5 e 2", func(t *testing.T) {
        resultado, err := Calculadora(5, 2, "-")
        verificarResultado(t, resultado, 3, err, false)
    })

    // Teste de Multiplicação
    t.Run("Multiplicação de 5 e 2", func(t *testing.T) {
        resultado, err := Calculadora(5, 2, "*")
        verificarResultado(t, resultado, 10, err, false)
    })

    // Teste de Divisão
    t.Run("Divisão de 5 por 2", func(t *testing.T) {
        resultado, err := Calculadora(5, 2, "/")
        verificarResultado(t, resultado, 2.5, err, false)
    })

    // Teste de Divisão por zero
    t.Run("Divisão de 5 por 0", func(t *testing.T) {
        resultado, err := Calculadora(5, 0, "/")
        verificarResultado(t, resultado, 0, err, true) // Espera erro
    })

    // Teste de Operador inválido
    t.Run("Operador inválido", func(t *testing.T) {
        resultado, err := Calculadora(5, 2, "%")
        verificarResultado(t, resultado, 0, err, true) // Espera erro
    })
}
