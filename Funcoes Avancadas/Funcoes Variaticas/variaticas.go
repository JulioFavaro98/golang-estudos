package main

import "fmt"

//Criando a função dessa maneira, eu não crio um limitador para a quantidade de dados inseridos na função
func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

//Dessa maneira ele irá printar apenas o texto unico que criamos, porém os números serão variáveis a quantidade de números que eu inserir na função
func escrevendo(texto string, numeros ...int) { //O parametro variático obrigatoriamente deve ser declarado por ultimo
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	totalDaSoma := soma(1, 23, 43, 22, 56, 78)
	fmt.Println(totalDaSoma)

	escrevendo("Olá Mundo", 1, 5, 6, 4, 8, 9, 99)

}
