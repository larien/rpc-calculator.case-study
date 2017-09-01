package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"strings"
)

func Register(server *rpc.Server, zig shared.Arith) {
	// Registra a interface Ziguifryda com o nome de "Ziguifryda"
	server.RegisterName("Ziguifryda", zig)
}

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

	// Cria uma instância da estrutura que implementa a interface Ziguifryda
	zig := new(Ziguifryda)

	// Registra um novo servidor RPC e registra a estrutura criada acima
	server := rpc.NewServer()

	// Registra a estrutura criada acima
	Register(server, zig)

	// registers an HTTP handler for RPC messages on rpcPath, and a debugging handler on debugPath
	server.HandleHTTP("/", "/debug")

	// Espera receber pacotes de dados via TCP na porta especificada
	l, e := net.Listen("tcp", ":8989")
	if e != nil {
		log.Fatal("Erro:", e)
	}

	// Inicia o servidor HTTP do Go no socket especificado por 1
	http.Serve(l, nil)

	fmt.Print("Insira a expressão infixa: ")
	infixString, err := ReadFromInput()
	if err != nil {
		fmt.Println("Falha ao receber expressão:", err.Error())
		return
	}

	fmt.Println("Notação pós-fixa:", compute.Convert(infixString))
	return
}
