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
)
var now=time.Now()

func init() {
	// runtime.NumCPU()
	runtime.GOMAXPROCS(8)
	fmt.Println("init:",int(time.Now().Sub(now).Seconds()))
	w:=make(chan bool)
	go func () {
		time.Sleep(time.Second*3)
		w<-true
	}()
	<-w
}
func main() {
	fmt.Println("main:",int(time.Now().Sub(now).Seconds()))
}

