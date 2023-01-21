package main

import "fmt"

func funcao1() {
	fmt.Println("Executando função 1")
}

func funcao2() {
	fmt.Println("Executando função 2")
}

func calculoMedia(n1, n2 float32) bool {
	defer fmt.Println("Média calculada!! Retornando o resultado!!") //Clausula Defer irá adiar a execução da função até o último momento possível
	fmt.Println("Entrando na Função para verificar se o aluno foi APROVADO")
	media := n1 + n2/2
	if media >= 6 {
		return true
	}

	return false
}

func main() {
	defer funcao1() //Clausula Defer irá adiar a execução da função até o último momento possível
	funcao2()
	fmt.Println(calculoMedia(4, 5))

}
