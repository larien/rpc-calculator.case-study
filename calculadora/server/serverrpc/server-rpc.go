package serverrpc

import (
    "calculadora/server/compute"
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

func (t *Ziguifryda) Function(args *common.Func, reply *common.Reply) error {

    funcao := compute.ConvertFunction(args.Funcao)

    switch funcao{
        case 1:
        // 1 - Seno
        reply.Resp := math.Sin(args.Value)
        reply.Success = true
        // 2- Cosseno
        case 2:
        reply.Resp := math.Cos(args.Value)
        reply.Success = true
        // 3 - Tangente - precisa de complex128
        case 3:
        reply.Resp := 0
        reply.Success = false
        // 4 - Cotangente - precisa de complex128
        case 4:
        reply.Resp := 0
        reply.Success = false
        // 5 - Seno inverso
        case 5:
        reply.Resp := math.Asen(args.Value)
        reply.Success = true
        // 6 - Cosseno inverso
        case 6:
        reply.Resp := math.Acos(args.Value)
        reply.Success = true
        // 7 - Tangente inversa - precisa de complex128
        case 7:
        reply.Resp := 0
        reply.Success = false
        // 8 - Cotangente inversa - precisa de complex128
        case 8:
        reply.Resp := 0
        reply.Success = false
        // 9 - Log
        case 9:
        reply.Resp := math.Log(args.Value)
        reply.Success = true
        // 10 - Ln
        case 10:
        reply.Resp := math.Log1p(args.Value)
        reply.Success = true
        //  11 - Lg
        case 11:
        reply.Resp := math.Log10(args.Value)
        reply.Success = true
        // 12 - Mod
        case 12:
        if args.Value < 0 {
            reply.Resp := args.Value*-1
        reply.Success = true
        }
        reply.Resp := args.Value
        reply.Success = true
        // 13 - Raiz quadrada
        case 13:
        reply.Resp := math.Sqrt(args.Value)
        reply.Success = true
        // 14 - Abs
        case 14:
        reply.Resp := math.Abs(args.Value)
        reply.Success = true
        default:
        reply.Resp := 0
        reply.Success = false
    }
    return nil
}