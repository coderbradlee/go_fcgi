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
func test(s string,n ...int)string {
	var x int
	for _,i:=range n{
		x+=i
	}
	return fmt.Sprintf(s,x)
}
func main() {
	
	a:=[3]int{0,1,2}
	for i,v:=range a{
		if i==0{
			a[1],a[2]=998,999
			fmt.Println(a)
		}
		a[i]=v+100
	}
	fmt.Println(a)
	fmt.Println(test("sum:%d",a...))
}

