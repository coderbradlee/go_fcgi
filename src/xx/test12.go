package main

import (
	"fmt"
	// "regexp"
	// "os"
	// "bufio"
	// "runtime"
	// "io"
	// "sync"
)
func test_defer(a int) {
	defer fmt.Println("a")
	defer fmt.Println("b")
	defer func () {
		fmt.Println(100/a)
	}()
	defer fmt.Println("c")
}
func main() {
	test_defer(0)
}

