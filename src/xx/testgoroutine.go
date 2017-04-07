package main

import (
	"fmt"
	"os"
	"runtime"
)

func init() {
	// runtime.NumCPU()
	runtime.GOMAXPROCS(runtime.NumCPU())
}
var goal int
func primetask(c chan int) {
	p:=<-c
	if p>goal{
		os.Exit(0)
	}
	fmt.Println(p)
	nc:=make(chan int)
	go primetask(nc)
	for{
		i:=<-c
		if i%p!=0{
			nc<-i
		}
	}
}
func main() {
	goal=100
	c:=make(chan int)
	go primetask(c)
	for i:=2;;i++{
		c<-i
	}
}
