package main

import "fmt"

func main() {
	fmt.Println("Trabalhando com loop")
	var frutas [5]string

	//colocando dados no array
	for i := 0; i < 5; i++ {
		fmt.Printf("Digite uma  fruta para a lista")
		fmt.Scanln(&frutas[i])
	}

	//imprimindo array
	for i := 0; i < len(frutas); i++ {
		fmt.Println(frutas[i])
	}

	//loop tipo while
	i := 0
	for i < 100 {
		fmt.Println(i)
		i++
	}

	//loop range
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	for i, num := range nums {
		fmt.Println(i, num)
	}

	//loop aninhados
	var x int = 5
	var y int = 5
	for i := 0; i < x; i++ {
		fmt.Println("oi")
		for z := 0; z < y; z++ {
			fmt.Printf("hello")
		}
	}
}
