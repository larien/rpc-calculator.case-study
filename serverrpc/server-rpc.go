package serverrpc

import (
    "compute"
    "common"
    "log"
)
type Ziguifryda int

func (t *Ziguifryda) Calculate(args *common.Args, reply *common.Reply) error {
    
	expressao := args.Expression

	// Verifica se a expressão é válida
	if !compute.Analyze(expressao) {
        reply.Resp = 5
        reply.Success = false
        log.Println(reply)
	} else {
	// Converte a expressão para pós-fixa
	posfixa := compute.Convert(expressao)

	// Calcula a conta e armazena o resultado na resposta
    reply.Resp = compute.Calculate(posfixa)
    reply.Success = true
    log.Println(reply)
    }
return nil
}
