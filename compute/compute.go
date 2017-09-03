package compute

import (
	"strings"
	"fmt"
	"log"
)

// Variável auxiliar para mostrar os logs na tela
var print string

// FIXME - A função não funciona corretamente - aceita funções inválidas
// Analyze é a função que verifica se a expressão inserida é válida ou não
func Analyze(expr string) bool {
	// Inicializa pilha
	var stack Stack

	print = fmt.Sprintf("Analizando expressão %s",expr)
	log.Println(print)

	// // Recebe o tamanho da expressão
	// tamanhoExpr := len(expr)

	// Entra no loop que varre a expressão
	for _, r := range expr {
		c := fmt.Sprintf("%c",r)
		log.Println(c, r)
		if c == "(" || c == "[" || c == "{" {
			stack.Push(c)
			log.Println("Entrou na entrada")
		} else if c == ")" || c == "]" || c == "}" {
			log.Println("Entrou na saída")
			if stack.Empty() {
				log.Println("A pilha está vazia antes do esperado")
				return false
			}
			p := fmt.Sprintf("%c",stack.Pop())
			// print = fmt.Sprintf("P = %d, C = %d",p,c)
			// log.Println(print)
			log.Println("P: ",p," C:", c)
			if p == "(" && c != ")" {
				log.Println("Há um caracter ')' sem inicializador")
				return false
			} else if p == "[" && c != "]" {
				log.Println("Há um caracter ']' sem inicializador")
				return false
			} else if p == "{" && c != "}" {
				log.Println("Há um caracter '}' sem inicializador")
				return false
			}
		} else {
			log.Println("Entrou em nada")
			return false
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

	print = fmt.Sprintf("Convertendo expressão %s para pós-fixa",expr)
	log.Println(print)

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

	print = fmt.Sprintf("A expressão %s foi convertida para pós-fixa: %s",expr, postfix)
	log.Println(print)
	// Retorna a expressão na notação pós-fixa sem espaços
	return strings.TrimSpace(postfix)
}

// Calculate calcula o resultado da expressão pós-fixa
// TODO - criar função calculadora
func Calculate(posfix string) float32 {
	//var stack Stack

	print = fmt.Sprintf("Calculando expressão %s",posfix)
	log.Println(print)
	
	return 2.0
}

// replace substitui os caracteres '[', ']', '{' e '}' por '(' e ')'
// FIXME - função desnecessária para a apresentação
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
