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
type Stringer interface{
	String() string
}
type Printer interface{
	Stringer
	Print()
}
type User struct{
	id int
	name string
}
func (self *User)String() string{
	return fmt.Sprintf("user %d,%s",self.id,self.name)
}
func (self *User)Print() {
	fmt.Println(self.String())
}
func Print(v interface{}) {
	fmt.Printf("%T:%v\n",v,v)
}
func main() {
	Print(1)
	Print("hello world!")
}

