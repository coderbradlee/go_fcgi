package main

import (
	"fmt"
	// "regexp"
	// "os"
	// "bufio"
	"runtime"
	// "io"
	// "sync"
	// "testing"
	// "math"
	"time"
	"unsafe"
)
type User struct{
	name string
}
type Admin struct{
	User
	title string
}
func init() {
	// runtime.NumCPU()
	runtime.GOMAXPROCS(runtime.NumCPU())
}
func main() {
	var u Admin
	t:=reflect.TypeOf(u)
	for i,n:=0,t.NumField();i<n;i++{
		f:=t.Field(i)
		fmt.Println(f.Name,f.Type)
	}
}

