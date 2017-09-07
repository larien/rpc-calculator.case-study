package common

// Args contém a expressão inserida pelo usuário
type Args struct {
	Expression string
}

type Reply struct {
	Resp float64
	Success bool
}

type Func struct {
	Funcao string
	Valor float64
}