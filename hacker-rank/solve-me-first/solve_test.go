package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveMeFirst(t *testing.T) {
	testCases := []struct {
		desc     string
		x, y     uint32
		expected uint32
	}{
		{
			desc: "sample test case 0",
			x:    2, y: 3,
			expected: 5,
		},
		{
			desc: "sample test case 1",
			x:    100, y: 1_000,
			expected: 1_100,
		},
		{
			desc: "lower bound",
			x:    1, y: 1,
			expected: 2,
		},
		{
			desc: "upper bound",
			x:    1_000, y: 1_000,
			expected: 2_000,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			assert.Equal(t, tc.expected, solveMeFirst(tc.x, tc.y))
		})
	}
}
