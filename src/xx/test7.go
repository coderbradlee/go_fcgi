package main

import (
	"fmt"
	"strings"
)
func print_s(key,value interface{}){
	fmt.Printf("%v->%v\n",key,value)
}
func main() {
	var omap map
	words:=[]string{"a","B","c","D"}
	ww:=omap.NewCaseFoldedKeyed()
	for _,w :=range words{
		ww.Insert(w,strings.ToUpper(w))
	}
	ww.Do(print_s)
}
