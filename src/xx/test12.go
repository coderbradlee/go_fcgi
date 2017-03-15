package main

import (
	"fmt"
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
func benchmarkTest(b *testing.B) {
	for i:=0;i<b.N;i++{
		test()
	}
}
func benchmarkTestDefer(b *testing.B) {
	for i:=0;i<b.N;i++{
		test_defer()
	}
}
func main() {
	test_defer(0)
}

