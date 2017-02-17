 package main
 import (
    "log"
    "fmt"
    _"encoding/json"
    "net/http"
    "io/ioutil"
)
func poHandler (w http.ResponseWriter, r *http.Request) {
	if r.Method !="POST"{
		fmt.Fprint(w, "this interface should be post!")
	} else{
		body, _ := ioutil.ReadAll(r.Body)
		log.Printf(body)
		// decoder := json.NewDecoder(r.Body)
	 //    var t test_struct   
	 //    err := decoder.Decode(&t)
	 //    if err != nil {
	 //        panic(err)
	 //    }
	    defer r.Body.Close()
	    // log.Println(t.Test)
	    ret := `{
				   "error_code":"200",
				   "error_msg":"Goods received successfully at 2017-03-17 12:00:00",
				   "data":{
				      "goods_receipt_no":"GR-FR-20170226-000196",
				      "bill_type":"Goods Receipt",
				      "receive_by":"Enie Yang",
				      "company":"ReneSola France",
				      "receive_at":"2017-03-17 12:00:00"
				   },
				   "reply_time":"2017-03-17 12:00:00"
			   }`

		fmt.Fprint(w, ret)
	}

} 
