// Copyright 2015 @lamphanviet. All rights reserved.

// Package primes provides functions to work on prime numbers
package primes

import (
	"math/rand"
	"../../commons"
)

// Return list of primes to n and array of bool
// Complexity: O(nloglogn)
func Eratosthenes(n int) (isPrime []bool, primes []int) {
	isPrime = make([]bool, n)
	for i := 2; i < n; i++ {
		isPrime[i] = true
	}
	for i := 4; i < n; i += 2 {
		isPrime[i] = false
	}
	for i := 3; i * i < n; i += 2 {
		if isPrime[i] {
			for j := i * i; j < n; j += i + i {
				isPrime[j] = false
			}
		}
	}

	primes = make([]int, 0)
	primes = append(primes, 2)
	for i := 3; i < n; i += 2 {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}

	return isPrime, primes
}

func IsPrime(p int64, primes []int) {
	for i := 0; i < len(primes) && primes[i] * primes[i] <= p; i++ {
		if p % primes[i] == 0 {
			return false
		}
	}
	return true
}

func IsPrimeSqrt(p int64) {
	if p < 2 {
		return false
	}
	if p == 2 {
		return true
	}
	if p % 2 == 0 {
		return false
	}
	for i := 3; i * i <= p; i += 2 {
		if p % i == 0 {
			return false
		}
	}
	return true
}

// testing a prime number using MillerRabin algorithm
func IsPrimeMiller(p int64, iteration int) bool {
	if p < 2 {
		return false;
	}
	if p != 2 && p % 2 == 0 {
		return false;
	}
	s := p - 1;
	for (s % 2 == 0) {
		s /= 2;
	}
	for i := 0; i < iteration; i++ {
		a, temp := rand.Int() % (p - 1) + 1, s
		mod := commons.PowModulo(a, temp, p);
		for temp != p - 1 && mod != 1 && mod != p - 1 {
		mod = commons.PowModulo(mod, mod, p);
		temp *= 2;
		}
		if mod != p - 1 && temp % 2 == 0 {
			return false;
		}
	}
	return true;
}

