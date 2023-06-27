package main

import (
	"fmt"
	"strings"
)

func contarPalavras(s string) map[string]int {
	palavras := strings.Fields(s) // Divide a string em palavras
	ocorrencias := make(map[string]int)

	for _, palavra := range palavras {
		ocorrencias[palavra]++
	}

	return ocorrencias
}

func main() {
	texto := "O rato roeu a roupa do rei de Roma"
	resultado := contarPalavras(texto)

	for palavra, ocorrencias := range resultado {
		fmt.Printf("%s: %d\n", palavra, ocorrencias)
	}
}
