// Copyright 2015 @lamphanviet. All rights reserved.

// Package primes provides functions to work on prime numbers
package primes

import (
	"math/rand"

	"github.com/lamphanviet/goalgo/commons"
)

const (
	MillerRabinIteration = 50
)

// Eratosthenes returns list of primes number in range [2, n - 1]
func Eratosthenes(n int64) (primes []int64, isPrime []bool) {
	isPrime = make([]bool, n)
	for i := int64(2); i < n; i++ {
		isPrime[i] = true
	}
	for i := int64(4); i < n; i += 2 {
		isPrime[i] = false
	}
	for i := int64(3); i*i < n; i += 2 {
		if isPrime[i] {
			for j := i * i; j < n; j += i + i {
				isPrime[j] = false
			}
		}
	}

	primes = make([]int64, 0)
	primes = append(primes, 2)
	for i := int64(3); i < n; i += 2 {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}

	return primes, isPrime
}

// IsPrime tests a prime number by going through list of primes
func IsPrime(p int64, primes []int64) bool {
	for i := 0; i < len(primes) && primes[i]*primes[i] <= p; i++ {
		if p%primes[i] == 0 {
			return false
		}
	}
	return true
}

// IsPrimeSqrt tests a prime number with complexity is O(sqrt(p))
func IsPrimeSqrt(p int64) bool {
	if p < 2 {
		return false
	}
	if p == 2 {
		return true
	}
	if p%2 == 0 {
		return false
	}
	for i := int64(3); i*i <= p; i += 2 {
		if p%i == 0 {
			return false
		}
	}
	return true
}

// IsPrimeMiller tests a prime number using MillerRabin algorithm
func IsPrimeMiller(p int64, iteration int) bool {
	if p < 2 {
		return false
	}
	if p != 2 && p%2 == 0 {
		return false
	}
	s := p - 1
	for s%2 == 0 {
		s /= 2
	}
	for i := 0; i < iteration; i++ {
		a, temp := rand.Int63()%(p-1)+1, s
		mod := commons.PowModulo(a, temp, p)
		for temp != p-1 && mod != 1 && mod != p-1 {
			mod = commons.PowModulo(mod, mod, p)
			temp *= 2
		}
		if mod != p-1 && temp%2 == 0 {
			return false
		}
	}
	return true
}
