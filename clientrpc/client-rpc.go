package clientrpc

import (
	"log"
	"net/rpc"
	"calculadora/common"
)

// SendExpression é a função que envia a expressão inserida pelo usuário para o servidor calcular
func SendExpression(client *rpc.Client, exp string) (float64, error) {
	args := common.Args{Expression: exp}

	var reply common.Reply
	err := client.Call("Ziguifryda.Calculate", args, &reply)
	if err != nil {
		log.Fatal("Erro", err)
	}

	return reply.Resp, nil
}