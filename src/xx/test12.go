package main

import (
	"fmt"
	// "regexp"
	"os"
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
	a,b:=make(chan int,3),make(chan int)
	go func () {
		v,ok,s:=0,false,""
		for{
			select{
				case v,ok=<-a:s="a"
				case v,ok=<-b:s="b"
			}
			if ok{
				fmt.Println(s,v)
			}else{
				os.Exit(0)
			}
		}
	}()
	for i:=0;i<10;i++{
		select{
		case a<-i:
		case b<-i:
		}
	}
	close(a)
}

