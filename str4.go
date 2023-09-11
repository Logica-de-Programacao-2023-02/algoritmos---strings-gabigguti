//Peça ao usuário para digitar uma string e retorne a mesma string com as letras em maiúsculo.

package main

import (
	"fmt"
	"strings"
)

func main() {
	var s1 string
	fmt.Println("Digite uma string: ", s1)
	fmt.Scan(&s1)

	fmt.Print("A string em maiúsculo é: ", strings.ToUpper(s1))
}
