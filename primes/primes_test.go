// Copyright 2015 @lamphanviet. All rights reserved.

package primes_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/lamphanviet/goalgo/primes"
)

var (
	testPrimes = []int64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
)

func TestEratosthenes(t *testing.T) {
	primes, isPrime := Eratosthenes(3)
	assert.Equal(t, 1, len(primes))
	assert.Equal(t, testPrimes[0], primes[0])
	assert.Equal(t, 3, len(isPrime))
	assert.True(t, isPrime[2])

	primes, isPrime = Eratosthenes(200)
	primes = primes[:len(testPrimes)]
	assert.Equal(t, testPrimes, primes)
}

func TestIsPrime(t *testing.T) {
	notPrime := int64(0)
	for _, prime := range testPrimes {
		for notPrime != prime {
			assert.False(t, IsPrime(notPrime, testPrimes), "must not be prime %v", notPrime)
			assert.False(t, IsPrimeSqrt(notPrime), "must not be prime %v", notPrime)
			//assert.False(t, IsPrimeMiller(notPrime, MillerRabinIteration), "must not be prime %v", notPrime)
			notPrime++
		}
		assert.True(t, IsPrime(prime, testPrimes), "must be prime %v", prime)
		assert.True(t, IsPrimeSqrt(prime), "must be prime %v", prime)
		//assert.True(t, IsPrimeMiller(prime, MillerRabinIteration), "must be prime %v", prime)
		notPrime = prime + 1
	}
}
