// Copyright 2015 @lamphanviet. All rights reserved.
package commons

// Calculate base^exponent in O(log(exponent))
func Pow(base int64, exponent int64) int64 {
	res := int64(1)
	mul := base
	for exponent > 0 {
		if exponent&1 != 0 {
			res *= mul
		}
		mul *= mul
		exponent >>= 1
	}
	return res
}

// Calculate (base^exponent) % mod in O(log(exponent))
func PowModulo(base int64, exponent int64, mod int64) int64 {
	res := int64(1)
	mul := base % mod
	for exponent > 0 {
		if exponent&1 != 0 {
			res = (res * mul) % mod
		}
		mul = (mul * mul) % mod
		exponent >>= 1
	}
	return res % mod
}
