package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"Calculadora/compute"
)

// ReadFromInput lê a expressão do usuário - RETIRAR APÓS IMPLEMENTAÇÃO
func ReadFromInput() (string, error) {

	reader := bufio.NewReader(os.Stdin)
	s, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Falha na leitura da expressão")
	}

	return strings.TrimSpace(s), err
}

func main() {

	fmt.Print("Insira a expressão infixa: ")
	infixString, err := ReadFromInput()
	if err != nil {
		fmt.Println("Falha ao receber expressão:", err.Error())
		return
	}

	fmt.Println("Notação pós-fixa:", compute.Convert(infixString))
	return
}
