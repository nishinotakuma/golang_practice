package anagram

import (
	"testing"
)

func TestAnagram(t *testing.T) {
	tests := []struct {
		a, b string
		want bool
	}{
		{"aba", "baa", true},
		{"aaa", "baa", false}, // same characters but different frequencies
	}
	for _, test := range tests {
		got := anagram(test.a, test.b)
		if got != test.want {
			t.Errorf("[%q, %q], got %v, want %v",
				test.a, test.b, got, test.want)
		}
	}
}
