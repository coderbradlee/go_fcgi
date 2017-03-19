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
	"reflect"
	// "unsafe"
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
	u:=new(Admin)
	t:=reflect.TypeOf(u)
	if t.Kind()==reflect.Ptr{
		t=t.Elem()
	}
	for i,n:=0,t.NumField();i<n;i++{
		f:=t.Field(i)
		fmt.Println(f.Name,f.Type)
	}
}

