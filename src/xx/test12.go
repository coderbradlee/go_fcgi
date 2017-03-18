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
	"math"
)
func sum(id int) {
	var x int64
	for i:=0;i<math.MaxUint32;i++{
		x+=int64(i)
	}
	fmt.Println(id,x)
}
func init() {
	// runtime.NumCPU()
	runtime.GOMAXPROCS(8)
}
func main() {
	c:=make(chan int,3)
	var chan1 chan<- int=c
	var chan2 <-chan int=c
	chan1<-1
	xx:=<-chan2
	fmt.Println(xx)
}

