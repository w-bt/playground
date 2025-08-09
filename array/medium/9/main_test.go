package main

import (
	"testing"
)

func eqBoolSlice(a, b []bool) bool {
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

func countTrue(a []bool) int {
	n := 0
	for _, v := range a {
		if v {
			n++
		}
	}
	return n
}

func TestNewChristmasLights_ErrOnTooSmall(t *testing.T) {
	if _, err := NewChristmasLights(3); err == nil {
		t.Fatalf("expected error for lightCount < 4, got nil")
	}
}

func TestNextPattern_BounceAndEndsAlwaysOn(t *testing.T) {
	lc := 8
	cl, err := NewChristmasLights(lc)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectedMidIdx := []int{0, 1, 2, 3, 4, 5, 4, 3, 2, 1, 0}

	for step, midIdx := range expectedMidIdx {
		got := cl.Next()

		if len(got) != lc {
			t.Fatalf("step %d: expected length %d, got %d", step+1, lc, len(got))
		}

		if !got[0] || !got[lc-1] {
			t.Fatalf("step %d: ends must be true; got[0]=%v got[last]=%v", step+1, got[0], got[lc-1])
		}

		if ct := countTrue(got); ct != 3 {
			t.Fatalf("step %d: expected exactly 3 trues, got %d: %v", step+1, ct, got)
		}

		want := make([]bool, lc)
		want[0] = true
		want[lc-1] = true
		want[1+midIdx] = true

		if !eqBoolSlice(got, want) {
			t.Fatalf("step %d: unexpected state.\n got=%v\nwant=%v", step+1, got, want)
		}
	}
}

func TestLargeLength_EndsAndCount(t *testing.T) {
	lc := 12
	cl, _ := NewChristmasLights(lc)

	for i := 0; i < 20; i++ {
		state := cl.Next()
		if len(state) != lc {
			t.Fatalf("len mismatch: got %d want %d", len(state), lc)
		}
		if !state[0] || !state[lc-1] {
			t.Fatalf("ends not true on step %d: %v", i+1, state)
		}
		if ct := countTrue(state); ct != 3 {
			t.Fatalf("expected 3 trues, got %d on step %d: %v", ct, i+1, state)
		}
	}
}
