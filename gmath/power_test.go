// Copyright 2015 @lamphanviet. All rights reserved.

package gmath_test

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/lamphanviet/goalgo/gmath"
)

func TestPow(t *testing.T) {
	assert.Equal(t, int64(1), Pow(0, 0))
	assert.Equal(t, int64(0), Pow(0, 5))
	assert.Equal(t, int64(1), Pow(1, 0))
	assert.Equal(t, int64(1), Pow(1, 5))
	assert.Equal(t, int64(2), Pow(2, 1))
	assert.Equal(t, int64(32), Pow(2, 5))
	assert.Equal(t, int64(1024), Pow(2, 10))
	assert.Equal(t, int64(1024), Pow(2, 10))
	assert.Equal(t, int64(2176782336), Pow(6, 12))
}

func TestPowModulo(t *testing.T) {
	const mod = 1000000007
	// test correctness
	for i := 0; i < 10; i++ {
		base := rand.Int63() % mod
		exponent := rand.Int63() % 100000

		actual := PowModulo(base, exponent, mod)
		expected := int64(1)
		for j := int64(0); j < exponent; j++ {
			expected = (expected * base) % mod
		}

		assert.Equal(t, expected, actual)
	}

	// test speed: fail if this run forever!
	for i := 0; i < 1000000; i++ {
		PowModulo(rand.Int63(), rand.Int63(), mod)
	}
}
