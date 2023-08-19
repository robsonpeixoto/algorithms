package comparable

// Compare returns
//
//	-1 if x is less than y,
//	 0 if x equals y,
//	+1 if x is greater than y.
type Compare[T any] func(x, y T) int
