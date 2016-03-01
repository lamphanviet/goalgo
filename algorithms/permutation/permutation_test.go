// Copyright 2015 @lamphanviet. All rights reserved.
package permutation_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/lamphanviet/goalgo/algorithms/permutation"
)

var (
	testArrays = [][]int{
		{1, 2, 3, 3},
		{1, 3, 2, 3},
		{1, 3, 3, 2},
		{2, 1, 3, 3},
		{2, 3, 1, 3},
		{2, 3, 3, 1},
		{3, 1, 2, 3},
		{3, 1, 3, 2},
		{3, 2, 1, 3},
		{3, 2, 3, 1},
		{3, 3, 1, 2},
		{3, 3, 2, 1},
	}
)

func TestNextIntSlice(t *testing.T) {
	n := len(testArrays)
	a := make([]int, len(testArrays[0]))
	copy(a, testArrays[0])
	for i := 1; i <= n; i++ {
		ok := NextInts(a)
		if i == n {
			assert.False(t, ok)
			assert.Equal(t, testArrays[n-1], a)
		} else {
			assert.True(t, ok)
			assert.Equal(t, testArrays[i], a)
		}
	}
}

func TestPrevIntSlice(t *testing.T) {
	n := len(testArrays)
	a := make([]int, len(testArrays[n-1]))
	copy(a, testArrays[n-1])
	for i := n - 2; i >= -1; i-- {
		ok := PrevInts(a)
		if i < 0 {
			assert.False(t, ok)
			assert.Equal(t, testArrays[0], a)
		} else {
			assert.True(t, ok)
			assert.Equal(t, testArrays[i], a)
		}
	}
}
