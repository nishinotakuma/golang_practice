// go test -bench .
package main

import (
	"strings"
	"testing"
)

var args = []string{"1", "2", "3", "4"}

func BenchmarkEcho(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r, sep := "", ""
		for _, a := range args {
			r += sep + a
			sep = " "
		}
	}
}

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Join(args, " ")
	}
}
