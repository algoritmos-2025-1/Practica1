package queue

// Implementación de estructura Queue usando slices
type SliceQueue[T any] struct {
	items []T
}

// Crea una estructura Queue
func CreateSliceQueue[T any]() *SliceQueue[T] {
	return &SliceQueue[T]{
		items: make([]T, 0),
	}
}

// Encola un elemento
func (q *SliceQueue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

// Desencola un elemento
func (q *SliceQueue[T]) Dequeue() T {
	first := q.items[0]
	q.items = q.items[1:]
	return first
}

// Regresa el primer elemento
func (q *SliceQueue[T]) Peek() T {
	return q.items[0]
}

// Verifica si la cola es vacía
func (q *SliceQueue[T]) Empty() bool {
	return len(q.items) == 0
}

// Agrega un elemento
func (q *SliceQueue[T]) Add(item T) {
	q.Enqueue(item)
}

// Elimina un elemento
func (q *SliceQueue[T]) Remove() T {
	return q.Dequeue()
}

// Elige un elemento
func (q *SliceQueue[T]) Choose() T {
	return q.Peek()
}