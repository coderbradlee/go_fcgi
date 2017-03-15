package main

import (
	"testing"
)

func Benchmark_Test(b *testing.B) {
	for i:=0;i<b.N;i++{
		test()
	}
}
func Benchmark_TestDefer(b *testing.B) {
	for i:=0;i<b.N;i++{
		test_defer()
	}
}

