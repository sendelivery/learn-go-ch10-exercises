// Package main contains helper arithmetic functions that can be used to
// improve legibility when performing basic mathematical operations.
package main

import "golang.org/x/exp/constraints"

// Number represents any integer or floating point number.
type Number interface {
	constraints.Integer | constraints.Float
}

// Add takes in two [Number] values and returns the result of their addition.
//
// Click [here] to learn more about addition.
//
// [here]: https://www.mathisfun.com/numbers/addition.html
func Add[T Number](a, b T) T {
	return a + b
}
