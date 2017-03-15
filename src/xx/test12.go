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
type Manager struct{
	User
}
func (self *User)ToString()string {
	return fmt.Sprintf("user:%p,%v",self,self)
}
func (self *Manager)ToString()string {
	return fmt.Sprintf("user:%p,%v",self,self)
}
func test_override() {
	m:=Manager{User{1,"tom"}}
	fmt.Println(m.ToString())
	fmt.Println(m.User.ToString())
}
func main() {

}

