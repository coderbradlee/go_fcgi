package main
//#cgo CFLAGS: -I../wkhtmltox/include 
//#cgo LDFLAGS: -lwkhtmltox
//#include "pdf.h"
import "C"
import (
    _"log"
    "fmt"
    "net/http"
)

func pdfHandler (w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "pdf!")
  C.Convert()
} 
