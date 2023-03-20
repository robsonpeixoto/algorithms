package chap01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	testCases := []struct {
		desc        string
		list        []int
		item        int
		expectedIdx int
		expectedErr error
	}{
		{
			desc:        "first item",
			list:        []int{1, 2, 3, 4, 5},
			item:        1,
			expectedIdx: 0,
			expectedErr: nil,
		},
		{
			desc:        "last item",
			list:        []int{1, 2, 3, 4, 5},
			item:        5,
			expectedIdx: 4,
			expectedErr: nil,
		},
		{
			desc:        "on half",
			list:        []int{1, 2, 3, 4, 5},
			item:        3,
			expectedIdx: 2,
			expectedErr: nil,
		},
		{
			desc:        "not found - smaller than all items",
			list:        []int{1, 2, 3, 4, 5},
			item:        0,
			expectedIdx: 0,
			expectedErr: ErrNotFound,
		},
		{
			desc:        "not found - bigger than all items",
			list:        []int{1, 2, 3, 4, 5},
			item:        6,
			expectedIdx: 0,
			expectedErr: ErrNotFound,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			idx, err := BinarySearch(tc.list, tc.item)
			assert.Equal(t, tc.expectedIdx, idx)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}
