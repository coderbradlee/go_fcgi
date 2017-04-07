package main
import (
	"fmt"
)
/*
#include <stdlib.h>
#include <stdio.h>
void print(char* cstr,int i) 
{
	sprintf(cstr,"%d\n",i);
    printf("%s", cstr);
}
*/
import "C"
import "unsafe"
// C语言的函数可以通过C这个包来访问
// 所有C语言的函数都可以视为是一个多值返回函数，第二个返回值是错误码 errno
// Golang不支持调用C里面的宏(例如printf用到宏)，这时候需要用C语言写个函数包装一下。

func random()int {
	return int(C.random())
}
func seed(i int) {
	C.srandom(C.uint(i))
}

func main() {
	seed(100)
	fmt.Println(random())
	var gostr string="1234"
	cstr:=C.CString(gostr)
	defer C.free(unsafe.Pointer(cstr))
	// 
	fmt.Println(gostr)
	C.print(cstr,123)
	// 
}