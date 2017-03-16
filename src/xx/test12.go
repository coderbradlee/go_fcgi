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
type Tester struct{
	s interface{
		String() string
	}
}
type User struct{
	id int
	name string
}
func (self *User)String() string{
	return fmt.Sprintf("user %d,%s",self.id,self.name)
}
func main() {
	u:=User{1,"tom"}
	t:=Tester{&u}
	u.id=2
	u.name="xx"
	fmt.Println(t.s.String())
}

