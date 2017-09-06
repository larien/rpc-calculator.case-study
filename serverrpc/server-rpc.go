package serverrpc

import (
    "calculadora/compute"
    "calculadora/common"
    "log"
)
type Ziguifryda int

func (t *Ziguifryda) Calculate(args *common.Args, reply *common.Reply) error {
    
	expressao := args.Expression

    // Verifica se a expressão é válida
    ok, err := compute.Analyze(expressao)
    if err != nil{
        log.Fatal("Falha na análise: "+err.Error())
    }
	if  !ok{
        reply.Resp = 5
        reply.Success = false
        log.Println(reply)
	} else {
        
	// Converte a expressão para pós-fixa
	posfixa := compute.Convert(expressao)

	// Calcula a conta e armazena o resultado na resposta
    resp, err := compute.Calculate(posfixa)
    if err != nil{
        log.Fatal("Falha no cálculo:"+err.Error())
    }
    reply.Resp = resp
    reply.Success = true
    log.Println(reply)
    }
return nil
}
