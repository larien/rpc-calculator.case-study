package clientrpc

import (
	"log"
	"net/rpc"
)

// Ziguifryda é a estrutura que faz a conexão do client com o servidor RPC
type Ziguifryda struct {
	client *rpc.Client
}

// Multiply é a função de exemplo que faz a chamada para a função de multiplicação
func (t *Ziguifryda) Multiply(a, b int) int {
	args := &shared.Args{a, b}
	var reply int
	err := t.client.Call("Server-RPC.Multiply", args, &reply)
	if err != nil {
		log.Fatal("Erro", err)
	}
	return reply
}
