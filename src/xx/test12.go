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
type User struct{
	id int
	name string
}
func (self User)String()string {
	return fmt.Sprintf("%d,%s",self.id,self.name)
}
func main() {
	var empty interface{}=User{1,"tom"}
	switch v:=empty.(type){
		case nil:
			fmt.Println("nil")
		case fmt.Stringer:
			fmt.Println(v)
		case func()string:
			fmt.Println(v())
	}
}

