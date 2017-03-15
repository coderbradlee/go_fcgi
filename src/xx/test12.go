package main

import (
	"fmt"
	// "regexp"
	// "os"
	// "bufio"
	// "runtime"
	// "io"
	"sync"
	// "testing"
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
func test_panic() {
	defer func () {
		if err:=recover();err!=nil{
			fmt.Println(err.(string))
		}
	}()
	panic("panic error")
}
func main() {
	// test_defer(0)
}

