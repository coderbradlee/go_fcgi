package main

import (
	// "fmt"
	// "regexp"
	// "os"
	// "bufio"
	// "runtime"
	// "io"
	"sync"
	"testing"
)
var lock sync.Mutex
func test() {
	lock.Lock()
	lock.Unlock()
}
func test_defer() {
	lock.Lock()
	defer lock.Unlock()
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

