package compute

import (
	"strings"
	"fmt"
	"log"
	"math"
	"strconv"
	"reflect"
	"errors"
)

// Variável auxiliar para mostrar os logs na tela
var print string

// Analyze é a função que verifica se a expressão inserida é válida ou não
func Analyze(expr string) (bool, error) {

	// Inicializa pilha
	var stack Stack

	print = fmt.Sprintf("Analizando expressão %s",expr)
	log.Println(print)

	// // Recebe o tamanho da expressão
	// tamanhoExpr := len(expr)

	// Entra no loop que varre a expressão
	for _, r := range expr {
		c := string(r)
		print = fmt.Sprintf("Letra em análise: %s",c)
		if c == "(" || c == "[" || c == "{" {
			stack.Push(c)
		} else if c == ")" || c == "]" || c == "}" {
			print = fmt.Sprintf("Entrou na saída: %s",c)
			log.Println(print)
			if stack.Empty() {
				return false, errors.New("A pilha está vazia antes do esperado")
			}
			p := stack.Pop().(string)
			print = fmt.Sprintf("Comparando P = %s, C = %s",p,c)
			log.Println(print)
			if p == "(" && c != ")" {
				return false, errors.New("Há um caracter ')' sem inicializador")
			} else if p == "[" && c != "]" {
				return false, errors.New("Há um caracter ']' sem inicializador")
			} else if p == "{" && c != "}" {
				return false, errors.New("Há um caracter '}' sem inicializador")
			}
		}
	}
	return stack.Empty(),nil
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
			// Se o operando tiver prioridade menor, desempilha
			// tudo até achar o parênteses inicial
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

// IsNumeric verifica se uma parte de string é um número válido
func IsNumeric(s string) bool {
   _, err := strconv.ParseFloat(s, 64)
   return err == nil
}

// Calculate calcula o resultado da expressão pós-fixa
func Calculate(postfix string) (float64, error) {

	print = fmt.Sprintf("Calculando expressão %s",postfix)
	log.Println(print)

	var stack Stack

	for _, r := range postfix {
		// r vem como rune, então precisamos converter para caracter string
		c := string(r)
		print = fmt.Sprintf("Verificando: %v",c)
			log.Println(print)
			// Se for número, coloca na pilha
		if IsNumeric(c){
			stack.Push(c)
			// Ignora espaços em branco
	} else if c != " " {
			var a, b float64
			// Pela pilha retornar uma interface, temos que confirmar se ela é
			// float64 ou string e fazer os devidos tratamentos à variável	
			if v := reflect.TypeOf(stack.Top()); v.Kind() == reflect.String{
			b,_ = strconv.ParseFloat(stack.Pop().(string),64)
			} else if v := reflect.TypeOf(stack.Top()); v.Kind() == reflect.Float64{
			b,_ = stack.Pop().(float64)
			}
			if v := reflect.TypeOf(stack.Top()); v.Kind() == reflect.String{
			a,_ = strconv.ParseFloat(stack.Pop().(string),64)
			} else if v := reflect.TypeOf(stack.Top()); v.Kind() == reflect.Float64{
			a,_ = stack.Pop().(float64)
			}
			print = fmt.Sprintf("A: %.2f | B: %.2f",a,b)
			log.Println(print)
			if c == "+" {
				stack.Push(a+b)
			} else if c == "-" {
				stack.Push(a-b)
			} else if c == "*" {
				stack.Push(a*b)
			} else if c == "/" {
				stack.Push(a/b)
			} else if c == "^" {
				stack.Push(math.Pow(a, b))
			}
		}
	}
	// Retorna erro caso sobre mais de um valor na pilha
	if stack.size > 1{
		return 0, errors.New("Sobrou caracteres demais no final da conta")
	} else {
		// Mesmo caso relacionado à interface
		if v := reflect.TypeOf(stack.Top()); v.Kind() == reflect.String{
			resp,_ := strconv.ParseFloat(stack.Pop().(string),64)
			return resp, nil
			} else if v := reflect.TypeOf(stack.Top()); v.Kind() == reflect.Float64{
			resp,_ := stack.Pop().(float64)
			return resp, nil
			} else {
				// Retorna erro caso não seja tipo string nem float64
			return 0, errors.New("Valor tem tipo irreconhecível")
		}
	}
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