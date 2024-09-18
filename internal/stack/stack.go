package stack

// Define una pila genérica
type Stack[T any] interface {
	// Inserta un elemento en la pila
	Push(item T)

	// Retira un elemento de la pila
	Pop() T

	// Regresa el primer elemento
	Peek() T

	// Verifica si la pila es vacía
	Empty() bool
}