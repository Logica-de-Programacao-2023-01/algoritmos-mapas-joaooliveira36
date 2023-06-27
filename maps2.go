package main

import "fmt"

func elementos_iguais(mapa1 map[string]int, mapa2 map[string]int) map[string]int {
	novo_mapa := make(map[string]int)

	for chave1, valor1 := range mapa1 {
		novo_mapa[chave1] = valor1
	}
	for chave2, valor2 := range mapa2 {
		novo_mapa[chave2] = valor2
	}

	return novo_mapa
}

func main() {
	mapa1 := map[string]int{"a": 1, "b": 2}
	mapa2 := map[string]int{"b": 3, "c": 4}

	novoMapa := elementos_iguais(mapa1, mapa2)
	fmt.Print(novoMapa)
}
