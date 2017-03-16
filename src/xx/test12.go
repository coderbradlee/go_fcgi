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
func (self *User)String()string {
	return fmt.Sprintf("%d,%s",self.id,self.name)
}
func main() {
	var empty interface{}=&User{1,"tom"}
	if i,ok:=empty.(fmt.Stringer);ok{
		fmt.Println(i,ok)
	}
	fmt.Println(empty.(*User))
}

