package main

import (
	"fmt"
	// "regexp"
	// "os"
	// "bufio"
	"runtime"
	// "io"
	"sync"
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
	
	wg:=new(sync.WaitGroup)
	wg.Add(2)
	go func () {
		defer wg.Done()
		for i:=0;i<6;i++{
			fmt.Println(i)
			if i==3{
				runtime.Gosched()
			}
		}
	}()
	go func () {
		defer wg.Done()
		fmt.Println("heeeee")	
	}()
	wg.Wait()
}

