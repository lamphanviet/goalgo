// Copyright 2015 @lamphanviet. All rights reserved.

package primes_test

import (
	"testing"

	"github.com/lamphanviet/goalgo/algorithms/primes"
	"github.com/stretchr/testify/assert"
)

var (
	testPrimes = []int64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
)

func TestEratosthenes(t *testing.T) {
	primes, isPrime := primes.Eratosthenes(3)
	assert.Equal(t, 1, len(primes))
	assert.Equal(t, testPrimes[0], primes[0])
	assert.Equal(t, 3, len(isPrime))
	assert.True(t, isPrime[2])
}

func TestIsPrime(t *testing.T) {
	notPrime := int64(0)
	for _, prime := range testPrimes {
		for notPrime != prime {
			assert.False(t, primes.IsPrime(notPrime, testPrimes))
			assert.False(t, primes.IsPrimeSqrt(notPrime))
			assert.False(t, primes.IsPrimeMiller(notPrime, primes.MillerRabinIteration))
			notPrime++
		}
		assert.True(t, primes.IsPrime(prime, testPrimes))
		assert.True(t, primes.IsPrimeSqrt(prime))
		assert.True(t, primes.IsPrimeMiller(prime, primes.MillerRabinIteration))
		notPrime = prime + 1
	}
}
