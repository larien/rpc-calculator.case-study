package serverrpc

import "log"

type Args struct {
	Expression string
}

type Ziguifryda int

// Calculate efetua o cálculo da expressão recebida
func (t *Ziguifryda) Calculate(args *Args, reply *float32) error {

	expressao := args.Expression

	// Verifica se a expressão é válida
	if !compute.Analyze(expressao) {
		log.Fatal("Falha na análise de expressão")
		return
	}

	// Converte a expressão para pós-fixa
	posfixa := compute.Convert(expressao)

	// Calcula a conta e armazena o resultado na resposta
	*reply = compute.Calculate(posfixa)
	return nil
}
