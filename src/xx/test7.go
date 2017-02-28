package main

import (
	"fmt"
	"strings"
)

func main() {
	var omap Map
	words:=[]string{"a","B","c","D"}
	ww:=omap.NewCaseFoldedKeyed()
	for _,w :=range words{
		ww.Insert(w,strings.ToUpper(w))
	}
	ww.Do(func(key,value interface{}){
		fmt.Printf("%v->%v\n",key,value)
		})
}
