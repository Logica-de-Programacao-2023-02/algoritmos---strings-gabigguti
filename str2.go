//Peça ao usuário para digitar uma string e retorne o número de caracteres nessa string.

package main

import "fmt"

func main() {

	var s1 string
	fmt.Println("Digite uma string: ", s1)
	fmt.Scan(&s1)

	fmt.Printf("O numero de letras da string é: %d", len(s1))
}
