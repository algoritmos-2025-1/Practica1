package collection

// Define una colección genérica
type Collection[T any] interface {
	// Agrega un elemento
	Add(item T)

	// Elimina un elemento
	Remove() T

	// Elige un elemento
	Choose() T

	// Verifica si la colección es vacía
	Empty() bool
}