 package main
 import (
    "log"
    "fmt"
    "time"
    "strconv"
    "go_redis_cluster"
    "net/http"
)
func pdfHandler (w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "pdf!")
} 
