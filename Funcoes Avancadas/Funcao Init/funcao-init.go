package main

import "fmt"

var n int

func init() {
	fmt.Println("Executando função init") //A função init SEMPRE será executada antes da função main, independente da ordem em que ela esteja
	n = 10
}

func main() {
	fmt.Println("Executando função main")
	fmt.Println(n)
}
