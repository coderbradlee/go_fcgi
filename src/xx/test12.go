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
	data:=make(chan int)
	exit:=make(chan bool)
	go func () {
		for d:=range data{
			fmt.Println(d)
		}
		fmt.Println("recv over")
		exit<-true
	}()
	data<-1
	data<-2
	data<-3
	close(data)
	fmt.Println("send over")
	<-exit
}

