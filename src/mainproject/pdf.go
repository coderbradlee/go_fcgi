package main
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
  C.Convert()
} 
