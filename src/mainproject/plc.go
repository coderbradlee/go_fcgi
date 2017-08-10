 package main
 import (
    "logger"
    "fmt"
    _"encoding/json"
    "net/http"
    "io/ioutil"
    _"bytes"
    _"time"
    // "strings"
)
func reportHandler (w http.ResponseWriter, r *http.Request) {
	////////////////////////////////
	addr := r.Header.Get("X-Real-IP")
	if addr == "" {
		addr = r.Header.Get("X-Forwarded-For")
		if addr == "" {
			addr = r.RemoteAddr
		}
	}

	// log.Printf("Started %s %s for %s", r.Method, r.URL.Path, addr)

/////////////////////////////////////////////////////////////////
	if r.Method !="POST"{
		fmt.Fprint(w, "this interface should be post!")
	} else{
		// body, err := ioutil.ReadAll(r.Body)
		// if err != nil {
	 // 		log.Println("ioutil.ReadAll error", err) 
 	// 	}
 	// 	sbody :=string(body)
 		//var ret string
		// // log.Println(sbody)
		// log.Printf("Started %s %s for %s:%s", r.Method, r.URL.Path, addr,sbody)
		// decoder := json.NewDecoder(r.Body)
		body, _:= ioutil.ReadAll(r.Body)
	    
		log_str:=fmt.Sprintf("Started %s %s for %s:%s\nresponse:%s", r.Method, r.URL.Path, addr,body,"")
	    logger.Info(log_str)


	    // log.Println(string(body))
	 //    var t DeliverGoodsForPO  
	 //    err_decode := json.Unmarshal(body, &t)
	 //    // bytes.Trim(body,"\\r\\n")
	 //    // line := strings.Trim(string(body), "\r\n")
		// defer r.Body.Close()
	     
	 //    // err_decode := decoder.Decode(&t)
	 //    if err_decode != nil {
	 //        // panic(err)
	 //        ret=`{"error_code":"`+error_json_decode+`","error_msg":`+err_decode.Error()+`,"data":{"reply_system":2},"reply_time":"`+time.Now().Format("2006-01-02 15:04:05")+`"}`
	 //        fmt.Fprint(w,ret )
	 //        // log.Printf("Started %s %s for %s:%s\nresponse:%s", r.Method, r.URL.Path, addr,body,ret)
	 //        log_str:=fmt.Sprintf("Started %s %s for %s:%s\nresponse:%s", r.Method, r.URL.Path, addr,body,ret)
	 //        logger.Info(log_str)
	 //        return;
	 //    }
	 //    // log.Println(t.Operation)
	 //    // var err_encode error
	 //    ret =get_response(&t)
	 //    // if err_encode != nil {
	 //    // 	// ret=`{"error_code":`+error_json_encode+`,"error_msg":`+err_encode.Error()+`,"data":{},"reply_time":"`+time.Now().Format("2006-01-02 15:04:05")+`"}`
	 //    //     // fmt.Fprint(w, ret)
	 //    //     // fmt.Println(ret)
	 //    //     // log.Fatal(err_encode.Error)
	 //    //     fmt.Fprint(w,ret )
	 //    //     log.Printf("Started %s %s for %s:%s\nrespose:%s\nerror:%s", r.Method, r.URL.Path, addr,"sbody",ret,err_encode.Error)
	 //    //     return;
	 //    // }
	 //    fmt.Fprint(w,ret )
	 //    // log.Printf("Started %s %s for %s:%s\nresponse:%s", r.Method, r.URL.Path, addr,body,ret)
	 //    log_str:=fmt.Sprintf("Started %s %s for %s:%s\nresponse:%s", r.Method, r.URL.Path, addr,body,ret)
	 //        logger.Info(log_str)
	}

} 
