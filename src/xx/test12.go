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
// func (self User)String()string {
// 	return fmt.Sprintf("%d,%s",self.id,self.name)
// }
func main() {
	var _ fmt.Stringer=(*User)(nil)
}

