package chap01

import (
	"errors"

	"github.com/robsonpeixoto/algorithms/comparable"
	"golang.org/x/exp/constraints"
)

var ErrNotFound = errors.New("item not found")

// BinarySearch find a item in a sorted list. Returns (0, ErrNotFound) if item
// does not exists.
func BinarySearch[T constraints.Ordered](list []T, item T) (int, error) {
	return BinarySearchBy(list, item, comparable.CompareOrdered[T])
}

func BinarySearchBy[T any](list []T, item T, compare comparable.CompareTo[T]) (int, error) {
	lowerBound := 0
	upperBound := len(list) - 1
	for lowerBound <= upperBound {
		middleIdx := (lowerBound + upperBound) / 2
		guess := list[middleIdx]
		switch cmp := compare(guess, item); {
		case cmp == 0:
			return middleIdx, nil
		case cmp > 0:
			upperBound = middleIdx - 1
		case cmp < 0:
			lowerBound = middleIdx + 1
		}
	}
	return 0, ErrNotFound
}
