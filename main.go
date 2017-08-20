package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadFromInput() (string, error) {

	reader := bufio.NewReader(os.Stdin)
	s, err := reader.ReadString('\n')

	return strings.TrimSpace(s), err
}

func main() {

	fmt.Print("Insira a expressão infixa: ")
	infixString, err := ReadFromInput()
	if err != nil {
		fmt.Println("Falha ao receber expressão:", err.Error())
		return
	}

	fmt.Println("Notação pós-fixa:", infixString)
	return
}
