 package main
 import (
    _"log"
    "fmt"
    
    "net/http"
)
func poHandler (w http.ResponseWriter, r *http.Request) {
	if r.Method !="POST"{
		fmt.Fprint(w, "this interface should be post!")
	} else{
		fmt.Fprint(w, "this is post!")
	}

} 
