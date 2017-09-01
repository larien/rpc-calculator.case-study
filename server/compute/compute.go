package compute

import "strings"

// Analyze é a função que verifica se a expressão inserida é válida ou não
func Analyze(expr string) bool {
	// Inicializa pilha
	var stack Stack

	// // Recebe o tamanho da expressão
	// tamanhoExpr := len(expr)

	// Entra no loop que varre a expressão
	for _, c := range expr {
		if c == '(' || c == '[' || c == '{' {
			stack.Push(c)
		} else if c == ')' || c == ']' || c == '}' {
			if stack.Empty() {
				return false
			}
			p := stack.Pop()

			if p == '(' && c != ')' {
				return false
			} else if p == '[' && c != ']' {
				return false
			} else if p == '{' && c != '}' {
				return false
			}
		}
	}
	return stack.Empty()
}

// IsOperator verifica se o caracter é um operador
func IsOperator(c uint8) bool {
	return strings.ContainsAny(string(c), "+ & - & * & / & ^")
}

// IsOperand verifica se o caracter é um operando
func IsOperand(c uint8) bool {
	return c >= '0' && c <= '9'
}

// GetPriority retorna o nível de prioridade do caracter
func GetPriority(op string) int {
	switch op {
	case "(":
		return 1
	case "+", "-":
		return 2
	case "*", "/":
		return 3
	case "^":
		return 4
	default:
		return 0
	}
}

// HasMorePriority verifica qual operador tem mais prioridade
func HasMorePriority(op1 string, op2 string) bool {
	pr1 := GetPriority(op1)
	pr2 := GetPriority(op2)
	return pr1 >= pr2
}

// Convert converte a expressão recebida para notação pós-fixa
func Convert(expr string) string {

	expr = replace(expr)
	var stack Stack

	postfix := ""

	length := len(expr)

	for i := 0; i < length; i++ {

		char := string(expr[i])
		// Ignora os espaços em branco
		if char == " " {
			continue
		}
		// Verifica os parênteses de entrada e saída
		if char == "(" {
			stack.Push(char)
		} else if char == ")" {
			for !stack.Empty() {
				str, _ := stack.Top().(string)
				if str == "(" {
					break
				}
				postfix += " " + str
				stack.Pop()
			}
			stack.Pop()
		} else if !IsOperator(expr[i]) {
			// Se o caracter não for um operador, é um operando
			j := i
			number := ""
			for ; j < length && IsOperand(expr[j]); j++ {
				number = number + string(expr[j])
			}
			postfix += " " + number
			i = j - 1
		} else {
			// Se o caracter for um operador, desempilha dois elementos,
			// faz a operação e empilha o resultado de volta
			for !stack.Empty() {
				top, _ := stack.Top().(string)
				if top == "(" || !HasMorePriority(top, char) {
					break
				}
				postfix += " " + top
				stack.Pop()
			}
			stack.Push(char)
		}
	}

	for !stack.Empty() {
		str, _ := stack.Pop().(string)
		postfix += " " + str
	}

	// Retorna a expressão na notação pós-fixa sem espaços
	return strings.TrimSpace(postfix)
}

func Calculate(posfix string) float32 {
	// TODO
}

// replace substitui os caracteres '[', ']', '{' e '}' por '(' e ')'
func replace(expr string) string {

	// Cria um novo replacer
	r := strings.NewReplacer(
		"{", "(",
		"[", "(",
		"}", ")",
		"]", ")")

	// Substitui todos os pares
	return r.Replace(expr)
}
