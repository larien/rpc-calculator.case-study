package main

import (
	"fmt"
	"log"
	"net/http"
	"net/rpc"
	//"shared" //Path to the package contains shared struct
)

func main() {

	// Tenta se conectar ao localhost:8989 usando o protocolo HTTP (a porta que o servidor RPC está ouvindo)
	client, err := rpc.DialHTTP("tcp", "localhost:8989")
	if err != nil {
		log.Fatal("Erro:", err)
	}

	// Cria roteador para servir à API para comunicação com front end
	router := NewRouter()

	// Executa servidor em localhost na porta 8080
	log.Fatal(http.ListenAndServe(":8080", router))

	// Cria uma estrutura que recebe todos os métodos.
	ziguifryda := &Ziguifryda{client: client}

	fmt.Println(ziguifryda.Multiply(5, 6))
}
