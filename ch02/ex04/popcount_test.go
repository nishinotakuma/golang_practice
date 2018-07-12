package popcount

import (
	"testing"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// loop
func PopCountLoop(x uint64) int {
	sum := 0
	for i := 0; i < 8; i++ {
		sum += int(pc[byte(x>>uint(i))])
	}
	return sum
}

// table
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// shift
func PopCountShift(x uint64) int {
	count := 0
	one := uint64(1)
	for i := 0; i < 64; i++ {
		if x&one > 0 {
			count++
		}
		x = x >> 1
	}
	return count
}

func bench(b *testing.B, f func(uint64) int) {
	for i := 0; i < b.N; i++ {
		f(uint64(i))
	}
}

func BenchmarkPopCount(b *testing.B) {
	bench(b, PopCount)
}

func BenchmarkLoop(b *testing.B) {
	bench(b, PopCountLoop)
}

func BenchmarkShift(b *testing.B) {
	bench(b, PopCountShift)
}
