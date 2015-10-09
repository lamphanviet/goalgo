// Copyright 2015 @lamphanviet. All rights reserved.

package permutation_test

import (
	"testing"
	. "../permutation"
)

func equalIntSlice(a, b []int) bool {
	if a == nil || b == nil {
		if a == nil && b == nil {
			return true
		}
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestNextIntSlice(t *testing.T) {
	data := []int{1, 2, 3}
	ok := NextInts(data)
	if !ok || !equalIntSlice(data, []int{1, 3, 2}) {
		t.Errorf("next permutation %v", data)
	}

	data = []int{1, 1, 2}
	ok = NextInts(data)
	if !ok || !equalIntSlice(data, []int{1, 2, 1}) {
		t.Errorf("next permutation %v", data)
	}

	data = []int{3, 2, 1}
	ok = NextInts(data)
	if ok || !equalIntSlice(data, []int{3, 2, 1}) {
		t.Errorf("next permutation %v", data)
	}
}

func TestPrevIntSlice(t *testing.T) {
	data := []int{3, 2, 1}
	ok := PrevInts(data)
	if !ok || !equalIntSlice(data, []int{3, 1, 2}) {
		t.Errorf("prev permutation %v", data)
	}

	data = []int{2, 1, 1}
	ok = PrevInts(data)
	if !ok || !equalIntSlice(data, []int{1, 2, 1}) {
		t.Errorf("prev permutation %v", data)
	}

	data = []int{1, 2, 3}
	ok = PrevInts(data)
	if ok || !equalIntSlice(data, []int{1, 2, 3}) {
		t.Errorf("prev permutation %v", data)
	}
}
