package main
//#cgo CFLAGS: -I/root/go_fcgi/src/wkhtmltox/include 
//#cgo LDFLAGS: -L/root/go_fcgi/src/wkhtmltox -lwkhtmltox
//#include "pdf.h"
import "C"
import (
    _"log"
    "fmt"
    "net/http"
)

func pdfHandler (w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "pdf!")
  C.Test(C.int(2))
} 
