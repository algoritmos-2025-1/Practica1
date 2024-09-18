package stack

// Implementación de estructura Stack usando slices
type SliceStack[T any] struct {
	items []T
}

// Crea una estructura Stack
func CreateSliceStack[T any]() *SliceStack[T] {
	return &SliceStack[T]{
		items: make([]T, 0),
	}
}

// Inserta un elemento en la pila
func (s *SliceStack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Retira un elemento de la pila
func (s *SliceStack[T]) Pop() T {
	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top
}

// Regresa el primer elemento
func (s *SliceStack[T]) Peek() T {
	return s.items[len(s.items)-1]
}

// Verifica si la pila es vacía
func (s *SliceStack[T]) Empty() bool {
	return len(s.items) == 0
}

// Agrega un elemento
func (q *SliceStack[T]) Add(item T) {
	q.Push(item)
}

// Elimina un elemento
func (q *SliceStack[T]) Remove() T {
	return q.Pop()
}

// Elige un elemento
func (q *SliceStack[T]) Choose() T {
	return q.Peek()
}