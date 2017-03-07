package main
/*
#include "stdio.h"

void test(int n) {
  char dummy[10240];

  printf("in c test func iterator %d\n", n);
  if(n <= 0) {
    return;
  }
  dummy[n] = '\a';
  test(n-1);
}
#cgo CFLAGS: -g
*/
//#cgo LDFLAGS: -lwkhtmltox
//#cgo CFLAGS: -I../wkhtmltox/include -L../wkhtmltox
import "C"
import (
    _"log"
    "fmt"
    "net/http"
)

func pdfHandler (w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "pdf!")
  C.test(C.int(2))
} 
