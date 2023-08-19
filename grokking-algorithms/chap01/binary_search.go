package chap01

import (
	"cmp"
	"errors"

	"github.com/robsonpeixoto/algorithms/comparable"
)

var ErrNotFound = errors.New("item not found")

// BinarySearch find a item in a sorted list. Returns (0, ErrNotFound) if item
// does not exists.
func BinarySearch[T cmp.Ordered](list []T, item T) (int, error) {
	return BinarySearchBy(list, item, cmp.Compare)
}

func BinarySearchBy[T any](list []T, item T, compare comparable.Compare[T]) (int, error) {
	lowerBound := 0
	upperBound := len(list) - 1
	for lowerBound <= upperBound {
		middleIdx := (lowerBound + upperBound) / 2
		guess := list[middleIdx]
		switch res := compare(guess, item); res {
		case 0:
			return middleIdx, nil
		case 1:
			upperBound = middleIdx - 1
		case -1:
			lowerBound = middleIdx + 1
		}
	}
	return 0, ErrNotFound
}
