package main
//#cgo CFLAGS: -Isrc/wkhtmltox/include 
//#cgo LDFLAGS: -Lsrc/wkhtmltox -lwkhtmltox
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
