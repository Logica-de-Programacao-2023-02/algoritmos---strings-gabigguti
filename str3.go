//Peça ao usuário para digitar uma string e um caractere e retorne o número de
//ocorrências desse caractere na string.

package main

import "fmt"

func main() {

	var s1, s2 string
	var count int

	fmt.Println("Digite uma string: ", s1)
	fmt.Scan(&s1)
	fmt.Print("Digite um caractere: ", s2)
	fmt.Scan(&s2)

	for _, c := range s1 {
		if string(c) == s2 {
			count++
		}
	}

	fmt.Printf("O numero de vezes do caractere é %d ", count)
}
