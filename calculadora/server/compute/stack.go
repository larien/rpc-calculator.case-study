package compute

// Stack é a estrutura da pilha
type Stack struct {
	top  *Element
	size int
}

// Element é a estrutura do elemento da pilha
type Element struct {
	value interface{}
	next  *Element
}

// Empty verifica se a pilha está vazia
func (s *Stack) Empty() bool {
	return s.size == 0
}

// Top retorna o elemento no topo da pilha
func (s *Stack) Top() interface{} {
	return s.top.value
}

// Push insere um elemento na pilha
func (s *Stack) Push(value interface{}) {
	s.top = &Element{value, s.top}
	s.size++
}

// Pop retira o elemento no topo da pilha
func (s *Stack) Pop() (value interface{}) {
	if s.size > 0 {
		value, s.top = s.top.value, s.top.next
		s.size--
		return value
	}
	return nil
}
