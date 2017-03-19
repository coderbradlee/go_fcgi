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
func (*User)ToString() {
	
}
func (Admin)ToString() {
	
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
	method:=func(t reflect.Type){
		for i,n:=0,t.NumMethod();i<n;i++{
			f:=t.Method(i)
			fmt.Println(f.Name,f.Type)
		}
	}
	method(reflect.TypeOf(u))
	fmt.Println("--------------------")
	method(reflect.TypeOf(*u))
}

