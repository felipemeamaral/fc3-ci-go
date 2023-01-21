package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(10, 10)

	if total != 20 {
		t.Errorf("Resultado da soma é inválido: Resultado %d. Esperado %d", total, 20)
	}
}