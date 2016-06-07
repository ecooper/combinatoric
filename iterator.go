// Package combinatoric is a simple Go port of the "combinatoric" parts of
// Python's itertools--specifically, combinations, permutations, and
// product.
//
// None of the iterators are threadsafe. Implement mutexes as required.
// Additionally, it should be assumed the return values for First and
// Next always share the same memory address. If return values must be
// persisted between iterations, copy them into another slice.
package combinatoric

// Iterator is the interface that wraps a basic iterator.
//
// Iterators are expected to track state and be able to calculate the
// number of iterations for a given implementation.
//
// First and Next should return nil when a result slice is not
// available.
type Iterator interface {
	First() []interface{}
	Next() []interface{}
	HasNext() bool
	Len() uint64
	Reset()
}
