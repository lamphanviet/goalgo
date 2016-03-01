// Copyright 2015 @lamphanviet. All rights reserved.

// Package permutation provides primitives for calculate permutations
// for slices and user-defined collections.
package permutation

// Interface defines methods of the permutation structure
type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// item[i] should be less than item[j]
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}

// Next produces next permutation and returns successful or not
func Next(data Interface) bool {
	dataLen := data.Len()
	i := dataLen - 1
	for i > 0 && !data.Less(i-1, i) {
		i--
	}

	if i <= 0 {
		return false
	}

	j := dataLen - 1
	for !data.Less(i-1, j) {
		j--
	}

	data.Swap(i-1, j)

	j = dataLen - 1
	for i < j {
		data.Swap(i, j)
		i++
		j--
	}

	return true
}

// Prev produces previous permutation and returns successful or not
func Prev(data Interface) bool {
	dataLen := data.Len()
	i := dataLen - 1
	for i > 0 && !data.Less(i, i-1) {
		i--
	}

	if i <= 0 {
		return false
	}

	j := dataLen - 1
	for !data.Less(j, i-1) {
		j--
	}

	data.Swap(i-1, j)

	j = dataLen - 1
	for i < j {
		data.Swap(i, j)
		i++
		j--
	}

	return true
}

// Convenience types for common cases

// IntSlice attaches the methods of Interface to []int
type IntSlice []int

func (p IntSlice) Len() int           { return len(p) }
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Int64Slice attaches the methods of Interface to []int64
type Int64Slice []int64

func (p Int64Slice) Len() int           { return len(p) }
func (p Int64Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Int64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Float64Slice attaches the methods of Interface to []float64
type Float64Slice []float64

func (p Float64Slice) Len() int           { return len(p) }
func (p Float64Slice) Less(i, j int) bool { return p[i] < p[j] || isNaN(p[i]) && !isNaN(p[j]) }
func (p Float64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// isNaN is a copy of math.IsNaN to avoid a dependency on the math package.
func isNaN(f float64) bool {
	return f != f
}

// StringSlice attaches the methods of Interface to []string
type StringSlice []string

func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// ByteSlice attaches the methods of Interface to []byte
type ByteSlice []byte

func (p ByteSlice) Len() int           { return len(p) }
func (p ByteSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p ByteSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func NextInts(a []int) bool { return Next(IntSlice(a)) }
func PrevInts(a []int) bool { return Prev(IntSlice(a)) }

func NextInts64(a []int64) bool { return Next(Int64Slice(a)) }
func PrevInts64(a []int64) bool { return Prev(Int64Slice(a)) }

func NextFloat64s(a []float64) bool { return Next(Float64Slice(a)) }
func PrevFloat64s(a []float64) bool { return Prev(Float64Slice(a)) }

func NextStrings(a []string) bool { return Next(StringSlice(a)) }
func PrevStrings(a []string) bool { return Prev(StringSlice(a)) }

func NextBytes(a []byte) bool { return Next(ByteSlice(a)) }
func PrevBytes(a []byte) bool { return Prev(ByteSlice(a)) }
