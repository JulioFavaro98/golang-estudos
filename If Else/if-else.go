package main

import "fmt"

func main() {
	fmt.Println("Estruturas de Controle")

	numero := 10
	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	if outroNumero := numero; outroNumero > 0 { //Desse modo eu inicializo uma variavel apenas para fazer a comparação, após isso eu não posso mais acessar pois ela é apagada
		fmt.Println("O número é maior que 0")
	} else if numero == 0 { //Else if é utilizado igual em C#
		fmt.Println("O número é igual a 0")
	} else {
		fmt.Println("O número é menor que 0")
	}
}
