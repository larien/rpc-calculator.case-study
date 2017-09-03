package common

// Args contém a expressão inserida pelo usuário
type Args struct {
	Expression string
}

type Reply struct {
	Resp float32
	Success bool
}