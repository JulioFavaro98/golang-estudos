package main

import (
	"fmt"
	"time"
)

func main() {
	func(texto string) {
		for i := 0; i < 10; i++ {
			fmt.Println(i+1, texto)
			time.Sleep(time.Second)
		}
	}("Olá Tete")
}
