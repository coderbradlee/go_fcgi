package main

import (
	"fmt"
	// "regexp"
	// "os"
	// "bufio"
	// "runtime"
	// "io"
	// "sync"
	// "testing"
)
type Tester interface{
	Do()
}
type FuncDo func()
func (self FuncDo)Do(){
	self()
}
func main() {
	var t Tester=FuncDo(func(){fmt.Println("hlll")})
	t.Do()
}

