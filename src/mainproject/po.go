 package main
 import (
    "log"
    "fmt"
    "encoding/json"
    "net/http"
)
func poHandler (w http.ResponseWriter, r *http.Request) {
	if r.Method !="POST"{
		fmt.Fprint(w, "this interface should be post!")
	} else{
		log.Printf(r.Body)
		// decoder := json.NewDecoder(r.Body)
	 //    var t test_struct   
	 //    err := decoder.Decode(&t)
	 //    if err != nil {
	 //        panic(err)
	 //    }
	    defer r.Body.Close()
	    // log.Println(t.Test)
		fmt.Fprint(w, "post ok!")
	}

} 
