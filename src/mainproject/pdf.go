package main
//#cgo CFLAGS: -g
//#cgo CFLAGS: -I../wkhtmltox/include -L../wkhtmltox
//#cgo LDFLAGS: -lwkhtmltox
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
