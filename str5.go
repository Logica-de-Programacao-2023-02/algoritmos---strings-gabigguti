//Peça ao usuário para digitar uma string e um número n e retorne a mesma string com as n primeiras letras em maiúsculo.

package main

import (
	"fmt"
	"strings"
)

func main() {

	var s1 string
	var n int

	fmt.Println("Digite uma string: ")
	fmt.Scan(&s1)
	fmt.Print("Digite um número: ")
	fmt.Scan(&n)

	fmt.Printf("A string é: %s ", strings.ToUpper(s1[:n])+s1[n:])
}
