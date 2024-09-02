package main

import (
	"fmt"
	"log"
)

func add(num1 float64, num2 float64) float64 {
	return num1 + num2
}
func sub(num1 float64, num2 float64) float64 {
	return num1 - num2
}
func mul(num1 float64, num2 float64) float64 {
	return num1 * num2
}
func div(num1 float64, num2 float64) float64 {
	return num1 / num2
}

func main() {
	fmt.Println("Aprendendo funçoes")
	fmt.Println("projeto 1 Calculadora basica")

	var simbolo string
	var num1 float64
	var num2 float64
	var res float64

	fmt.Println("Digite o primeiro numero")
	fmt.Scanln(&num1)
	fmt.Println("Digite a operaçao")
	fmt.Scanln(&simbolo)
	fmt.Println("Digite o segundo numero")
	fmt.Scanln(&num2)

	switch simbolo {
	case "+":
		res = add(num1, num2)
		fmt.Printf("O resultado da Soma: %.2f", res)
	case "-":
		res = sub(num1, num2)
		fmt.Printf("O resultado da Subtraçao: %.2f", res)
	case "*":
		res = mul(num1, num2)
		fmt.Printf("O resultado da Multiplicação: %.2f", res)
	case "/":
		res = div(num1, num2)
		fmt.Printf("O resultado da Divisão: %.2f", res)
	default:
		fmt.Println("simbolo nao encontrado")
		log.Panic("erro simbolo")
	}

}
