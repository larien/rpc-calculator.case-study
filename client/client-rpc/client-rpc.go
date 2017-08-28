package clientrpc

import (
	"log"
	"net/rpc"
)

type Args struct {
	Expression string
}

// Ziguifryda é a estrutura que faz a conexão do client com o servidor RPC
type Ziguifryda struct {
	client *rpc.Client
}

// SendExpression é a função que envia a expressão inserida pelo usuário para o servidor calcular
func (t *Ziguifryda) SendExpression(exp string) (float32, error) {
	args := Args{Expression}
	var reply float32
	err := t.client.Call("Server-RPC.Calculate", args, &reply)
	if err != nil {
		log.Fatal("Erro", err)
	}
	return 0, reply
}
