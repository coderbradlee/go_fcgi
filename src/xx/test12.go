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
func test_slice() {
	data:=[...]int{0,1,2,3,10:0}
	s:=data[:2:3]
	fmt.Println(&s[0],&data[0])
	fmt.Println(s,data)
	s=append(s,200)
	fmt.Println(&s[0],&data[0])
	fmt.Println(s,data)
	s=append(s,300,400)
	fmt.Println(&s[0],&data[0])
	fmt.Println(s,data)
}
func main() {
	// test_defer(0)
	test_slice()
}

