 package main
 import (
    "logger"
    "fmt"
    "encoding/json"
    "net/http"
    "io/ioutil"
    "bytes"
    "time"
    // "errors"
    // "runtime/pprof"
)
func deliver_goods_Handler (w http.ResponseWriter, r *http.Request) {
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
 		var ret string
		body, _:= ioutil.ReadAll(r.Body)
	    defer r.Body.Close()
	    var t DeliverGoodsForPO  
	    err_decode := json.Unmarshal(body, &t)
		
	    if err_decode != nil {
	        // panic(err)
	        ret=`{"error_code":"`+error_json_decode+`","error_msg":"`+err_decode.Error()+`","data":{"bill_no":"","bill_type":"Goods Receipt","receive_by":"",   "company":"","receive_at":""},"reply_time":"`+time.Now().Format("2006-01-02 15:04:05")+`"}`
	        fmt.Fprint(w,ret )
	        // log.Printf("Started %s %s for %s:%s\nresponse:%s", r.Method, r.URL.Path, addr,body,ret)
	        log_str:=fmt.Sprintf("Started %s %s for %s:%s response:%s", r.Method, r.URL.Path, addr,body,ret)
	        logger.Info(log_str)
	        return;
	    }
	    // log.Println(t.Operation)
	    // var err_encode error
	    ret =get_response_of_gdn(&t)
	    
	    fmt.Fprint(w,ret )
	    log_str:=fmt.Sprintf("Started %s %s for %s:%s response:%s", r.Method, r.URL.Path, addr,body,ret)
        logger.Info(log_str)
        // pprof.StopCPUProfile()
	}

} 

func get_response_of_gdn(t *DeliverGoodsForPO) (string){
	var sd=shared_data{}
	err_no,check_err:=gdn_check_data(t)
	if check_err!=nil{
		// return `{"error_code":"`+err_no+`","error_msg":"`+check_err.Error()+`","data":{"po_no":"`+t.Data.Purchase_order.Po_no+`","reply_system":2},"reply_time":"`+time.Now().Format("2006-01-02 15:04:05")+`"}`
		return `{"error_code":"`+err_no+`","error_msg":"`+check_err.Error()+`","data":{"bill_no":"`+t.Data.Deliver_notes[0].Gdn_no+`","bill_type":"Goods Delivery Note","receive_by":"",   "company":"","receive_at":"`+t.Request_time+`"},"reply_time":"`+time.Now().Format("2006-01-02 15:04:05")+`"}`
	}
	s,err:=insert_gdn_database(t,&sd)
	if err!=nil{
		// return `{"error_code":"`+s+`","error_msg":"`+err.Error()+`","data":{"po_no":"`+t.Data.Purchase_order.Po_no+`","reply_system":2},"reply_time":"`+time.Now().Format("2006-01-02 15:04:05")+`"}`
		return `{"error_code":"`+s+`","error_msg":"`+err.Error()+`","data":{"bill_no":"`+t.Data.Deliver_notes[0].Gdn_no+`","bill_type":"Goods Delivery Note","receive_by":"",   "company":"","receive_at":"`+t.Request_time+`"},"reply_time":"`+time.Now().Format("2006-01-02 15:04:05")+`"}`
	}
	
	json_ret:=&Response_json{Error_code:"200",Error_msg:"Goods received successfully at "+time.Now().Format("2006-01-02 15:04:05"),Data:Response_json_data{Bill_no:sd.goods_receipt_no,Bill_type:"Goods Delivery Note",Receive_by:"received",Company:"",Receive_at:time.Now().Format("2006-01-02 15:04:05")},Reply_time:time.Now().Format("2006-01-02 15:04:05")}
		
	var buffer bytes.Buffer
    enc := json.NewEncoder(&buffer)

    err_encode := enc.Encode(json_ret)
    if err_encode!=nil{
    	// return `{"error_code":"`+error_json_encode+`","error_msg":"`+err_encode.Error()+`","data":{"po_no":"`+t.Data.Purchase_order.Po_no+`","reply_system":2},"reply_time":"`+time.Now().Format("2006-01-02 15:04:05")+`"}`
    	return `{"error_code":"`+error_json_encode+`","error_msg":"`+err_encode.Error()+`","data":{"bill_no":"`+t.Data.Deliver_notes[0].Gdn_no+`","bill_type":"Goods Delivery Note","receive_by":"",   "company":"","receive_at":"`+t.Request_time+`"},"reply_time":"`+time.Now().Format("2006-01-02 15:04:05")+`"}`
    }
	return buffer.String()
}
func insert_gdn_database(t *DeliverGoodsForPO,sd *shared_data)(string,error){
    var level3_group errgroup
    var level4_group errgroup
    level3_group.Go(t,sd,insert_goods_delivery_note)
    // level3_group.Go(t,sd,insert_commercial_invoice)
    if s,err := level3_group.Wait(); err != nil {
     return s,err
    }else{
     // level4_group.Go(t,sd,insert_note_attachment)
     // level4_group.Go(t,sd,insert_note_detail)
     level4_group.Go(t,sd,insert_goods_receipt)
     s,err = level4_group.Wait()
     return s,err
    }
}