package main

import "fmt"

func soma_mapa(mapa map[string]int) int {
	soma := 0
	for _, valores := range mapa {
		soma += valores
	}

	return soma
}

func main() {
	map1 := map[string]int{"num1": 5, "num2": 2, "num3": 8, "num4": 3}
	soma := soma_mapa(map1)
	fmt.Print(soma)
}
