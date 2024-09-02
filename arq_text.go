package main

import (
	"fmt"
	"os"
)

func criar_txt(txt string) {
	fmt.Println("criando um arquivo txt")
	arquivo, err := os.Create("juramento_lanterna_verde.txt")
	if err != nil {
		fmt.Println("Erro ao criar arquivo", err)
		return
	}
	defer arquivo.Close()

	//escrever no arquivo
	_, err = arquivo.WriteString(txt)
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo", err)
		return
	}
	fmt.Println("Texto criado com sucesso")

}

func ler_txt() {
	conteudo, err := os.ReadFile("juramento_lanterna_verde.txt")
	if err != nil {
		fmt.Println("erro ao ler arquivo ", err)
		return
	}
	fmt.Println(string(conteudo))
}

func main() {
	//criando arq txt
	//lendo arquivo
	//escrevendo
	texto := `No dia mais claro,
Na noite mais densa,
Nenhum mal escapará à minha presença.
Todo aquele que venera o mal há de penar,
Quando o poder do Lanterna Verde enfrentar!`

	criar_txt(texto)
	ler_txt()

}
