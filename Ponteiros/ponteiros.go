package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	//COMO SE COMPORTAM DUAS VARIÁVEIS
	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	variavel2++
	fmt.Println(variavel1, variavel2)
	//_____________________________________________________________________
	//PONTEIRO É UMA REFERÊNCIA DE MEMÓRIA

	//COMO SE COMPORTAM UMA VARIÁVEL E UM PONTEIRO
	var variavel3 int = 100
	var ponteiro *int = &variavel3 //Para declarar um ponteiro é preciso colocar o sinal * na frente do tipo de dado como foi usado no *int e para referenciar a uma variável utilizamos o sinal & na frente do nome da variável

	fmt.Println(variavel3, *ponteiro) //O sinal de * na frente do ponteiro significa desreferenciação, caso não utilize isso ele apenas irá retornar o endereço da variável3

	variavel3 = 150
	fmt.Println(variavel3, *ponteiro)
}
