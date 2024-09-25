package main

import "fmt"

func main() {
	const ponto_ebulicao_kelvin float32 = 373
	result := ponto_ebulicao_kelvin - 273

	fmt.Printf("O ponto de ebulição da água em Celsius é: %g C \n", result)
}
