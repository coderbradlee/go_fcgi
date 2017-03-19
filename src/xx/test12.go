package main

import (
	"fmt"
	// "regexp"
	// "os"
	// "bufio"
	"runtime"
	// "io"
	// "sync"
	// "testing"
	// "math"
	"time"
	"unsafe"
)
type data struct{
	x [1024*100]byte
}
func test()unsafe.Pointer {
	p:=&data{}
	return unsafe.Pointer(p)
}
func init() {
	// runtime.NumCPU()
	runtime.GOMAXPROCS(runtime.NumCPU())
}
func main() {
	const N=100

	cache:=new([N]unsafe.Pointer)
	for i:=0;i<N;i++{
		cache[i]=test()
		time.Sleep(time.Millisecond)
	}
	fmt.Println("xxxxx")
}

