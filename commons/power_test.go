// Copyright 2015 @lamphanviet. All rights reserved.

package commons_test

import (
	"testing"
	. "../commons"
	"math/rand"
)

func TestPowModulo(t *testing.T) {
	const mod = 1000000007
	// test correctness
	for i := 0; i < 10; i++ {
		base := rand.Int63() % mod
		exponent := rand.Int63() % 100000

		result := PowModulo(base, exponent, mod)
		testResult := int64(1)
		for j := int64(0); j < exponent; j++ {
			testResult = (testResult * base) % mod
		}

		if result != testResult {
			t.Errorf("powModulo %v^%v = %v, correct is %v", base, exponent, result, testResult)
		}
	}

	// test speed: fail if this run forever!
	for i := 0; i < 1000000; i++ {
		PowModulo(rand.Int63(), rand.Int63(), mod)
	}
}