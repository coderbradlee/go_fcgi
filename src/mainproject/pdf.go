package main
//#cgo CFLAGS: -I../wkhtmltox/include 
//#cgo LDFLAGS: -L../wkhtmltox -lwkhtmltox
//void Test(int n);
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
