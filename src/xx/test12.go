package main

import (
	"fmt"
	"regexp"
	"os"
	"bufio"
	"runtime"
	"io"
	"sync"
)

func main() {
	
	a:=[3]int{0,1,2}
	for i,v:=range a{
		if i==0{
			a[1],a[2]=998,999
			fmt.Println(a)
		}
	}
}

