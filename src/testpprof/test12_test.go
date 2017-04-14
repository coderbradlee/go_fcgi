package main
import (
	"testing"
)
func BenchmarkPopcnt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := i
		x -= (x >> 1) & m1
		x = (x & m2) + ((x >> 2) & m2)
		x = (x + (x >> 4)) & m4
		_ = (x * h01) >> 56
	}
}
func BenckmarkFib(b *testing.B) {
	for n:=0;n<b.N;n++{
		fib(20)
	}
}
