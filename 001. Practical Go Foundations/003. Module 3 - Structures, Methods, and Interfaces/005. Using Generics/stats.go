package main

import (
	"errors"
	"fmt"
	"time"
)

// Named constraint.
type Number interface {
	~int | ~float64
}

// Generic data structure.
type Matrix[T Number] struct {
	Rows int
	Cols int
	data []T
}

func main() {
	// Generics example 1 - Generic function that handles multiple types.
	fmt.Println(Relu(7))
	fmt.Println(Relu(-1))
	fmt.Println(Relu(1.2))
	fmt.Println(Relu(time.February))
	fmt.Println()

	// Generics example 2 - Generic data structures.
	m, err := NewMatrix[float64](10, 3)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for i := range len(m.data) {
		m.data[i] = float64(i)
	}

	fmt.Println("m:", m)
	fmt.Println(m.At(0, 0))
	fmt.Println(m.At(1, 0))
	fmt.Println(m.At(1, 1))
	fmt.Println(m.At(3, 2))
	fmt.Println(m.At(9, 2))
	fmt.Println()

	// Exercise - a Max function for ints or floats (without using the built-in max function).
	fmt.Println(Max([]int{3, 1, 2}))     // 3, <nil>
	fmt.Println(Max([]float64{3, 1, 2})) // 3, <nil>
	fmt.Println(Max[int](nil))           // 0, Max of empty slice

}

/*
Generics example 1:

Use a generic function to handle multiple types.
Very useful if you have multiple, similar functions where the only differences between them are the types that they operate on.
Side note: T is a "type constraint" (not a new type).

func Relu[T ~int | ~float64](i T) T { // Function definition without a named constraint type.
*/
func Relu[T Number](i T) T { // Function definition with a named constraint type.
	if i < 0 {
		return 0
	}

	return i
}

/*
Generics example 2:
Generic data structures (with methods).
*/
func NewMatrix[T Number](rows, cols int) (*Matrix[T], error) {
	if rows <= 0 || cols <= 0 {
		return nil, fmt.Errorf("Bad dimensions: %d/%d", rows, cols)
	}

	m := Matrix[T]{
		Rows: rows,
		Cols: cols,
		data: make([]T, rows*cols),
	}

	return &m, nil

}

// At is a method on a generic data structure that return a specific value in a matrix, based on the row and column provided.
func (m *Matrix[T]) At(row, col int) T {
	i := (row * m.Cols) + col
	return m.data[i]
}

/*
Exercise: Write a Max function for ints or floats.
Don't use the built-in max function.
*/
func Max[T Number](values []T) (T, error) {
	if len(values) == 0 {
		var zero T
		return zero, errors.New("Max of an empty slice.")
	}

	m := values[0]
	for _, v := range values[1:] {
		if v > m {
			m = v
		}
	}

	return m, nil
}

/*
// Original, "non-generics" code.
func ReluInt(i int) int {
	if i < 0 {
		return 0
	}

	return i
}

func ReluFloat64(i float64) float64 {
	if i < 0 {
		return 0
	}

	return i
}
*/
