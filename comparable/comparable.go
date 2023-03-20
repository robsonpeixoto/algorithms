package comparable

import "golang.org/x/exp/constraints"

// CompareTo is a function that compare two elements of the same type. Must
// return a negativa, zero or positive number if the `lhs`(left hand side) is
// smaller, equals of bigger than the `rhs`(right hand size)
type CompareTo[T any] func(lhs, rhs T) int

func CompareOrdered[T constraints.Ordered](lhs, rhs T) int {
	ret := 0
	if lhs > rhs {
		ret = 1
	} else if lhs < rhs {
		ret = -1
	}
	return ret
}
