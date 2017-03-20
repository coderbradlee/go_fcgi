package main

import (
	"testing"
	"time"
)
func sum(n ...int)int {
	var c int
	for _,i:=range n{
		c+=i
	}
	return c
}
func TestSum(t *testing.T) {
	time.Sleep(time.Second*2)
	if sum(1,2,3)!=6{
		t.Fatal("sum error")
	}
}
func TestTimeout(t *testing.T) {
	time.Sleep(time.Second*5)
}

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

