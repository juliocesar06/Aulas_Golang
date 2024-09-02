package main

import "fmt"

func main() {

	//print
	fmt.Println("Hello World")

	//declaraçao de variaveis
	var a int = 30
	var b float32 = 3.99
	var er string = "JJJ"

	// print formatado
	fmt.Printf("int:%d float:%f  string:%s\n", a, b, er)

	//entrada de dados
	var nome string
	var salario float32

	fmt.Println("Digite seu nome:")
	fmt.Scanln(&nome)
	fmt.Println("O Meu Salrio é:")
	fmt.Scanln(&salario)

	fmt.Printf("salario: %.2f nome: %s\n", salario, nome)

	//condicionais simples (if,else)
	var nota1 float64
	var nota2 float64
	var nota3 float64
	var result float64

	fmt.Println("Digite a nota 1:")
	fmt.Scanln(&nota1)
	fmt.Println("Digite a nota 1:")
	fmt.Scanln(&nota2)
	fmt.Println("Digite a nota 1:")
	fmt.Scanln(&nota3)
	result = ((nota1 + nota2 + nota3) / 3) * 10
	if nota1 < 6 {
		fmt.Printf("Voçê tirou %.2f e nao passou\n", result)
		fmt.Printf("voce tirou %.1f %%", result)
	} else {
		fmt.Printf("Voçê tirou %.2f e passou\n", result)
		fmt.Printf("voce tirou %.1f %%", result)
	}

}
