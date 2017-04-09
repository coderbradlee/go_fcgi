 package main
 import (
    "logger"
    "fmt"
    "encoding/json"
    "net/http"
    "io/ioutil"
    "bytes"
    "time"
    "errors"
    // "runtime/pprof"
)
func po_check_Handler (w http.ResponseWriter, r *http.Request) {
	////////////////////////////////
	addr := r.Header.Get("X-Real-IP")
	if addr == "" {
		addr = r.Header.Get("X-Forwarded-For")
		if addr == "" {
			addr = r.RemoteAddr
		}
	}

/////////////////////////////////////////////////////////////////
	if r.Method !="POST"{
		fmt.Fprint(w, "this interface should be post!")
	} else{
 		var ret string
		body, _:= ioutil.ReadAll(r.Body)
	    
	    var t PoData  
	    err_decode := json.Unmarshal(body, &t)
		defer r.Body.Close()
	     
	    if err_decode != nil {
	        ret=`{"error_code":`+error_json_decode+`,"error_msg":"`+err_decode.Error()+`","data":{"bill_no":"","bill_type":"Purchase Order","receive_by":"",   "company":"","receive_at":""},"reply_time":"`+time.Now().Format("2006-01-02 15:04:05")+`"}`
	        fmt.Fprint(w,ret )
	        log_str:=fmt.Sprintf("Started %s %s for %s:%s \nresponse:%s", r.Method, r.URL.Path, addr,body,ret)
	        logger.Info(log_str)
	        return;
	    }
	    ret =get_check_response(&t)
	    
	    fmt.Fprint(w,ret )
	    log_str:=fmt.Sprintf("Started %s %s for %s:%s \nresponse:%s", r.Method, r.URL.Path, addr,body,ret)
        logger.Info(log_str)
	}

} 

func check_with_database(t *PoData,sd *shared_data,contact_account_id string)(string,error) {
	set_company_time_zone(t.Data.Purchase_order.Company,sd)

	var t_purchase_order purchase_order
		
	t_purchase_order.purchase_order_id=rand_string(20)
	t_purchase_order.po_no=t.Data.Purchase_order.Po_no
	t_purchase_order.po_date=t.Data.Purchase_order.Po_date//[0:11]

	t_purchase_order.status=t.Data.Purchase_order.Status

	//from t.Data.Purchase_order.Company find company_id
	company_id_chan :=make(chan string)
    go get_company_id_chan(company_id_chan,t.Data.Purchase_order.Company)
    t_purchase_order.company_id=<-company_id_chan
	vendor_basic_id_chan :=make(chan string)
    go get_vendor_basic_id_chan(vendor_basic_id_chan,t.Data.Purchase_order.Supplier)
    t_purchase_order.vendor_basic_id=<-vendor_basic_id_chan
	
	//待确定
	// t_purchase_order.contact_account_id=get_contact_account_id(t_purchase_order.company_id)
	
  	t_purchase_order.contact_account_id=contact_account_id

	t_purchase_order.payment_terms=t.Data.Purchase_order.Payment_terms
	t_purchase_order.requested_delivery_date=t.Data.Purchase_order.Requested_delivery_date//[0:11]
	// t.Data.Purchase_order.Ship_via select id
	// t_purchase_order.shipping_method_id=get_shipping_method_id(t.Data.Purchase_order.Ship_via)
	shipping_method_id_chan :=make(chan string)
    go get_shipping_method_id_chan(shipping_method_id_chan,t.Data.Purchase_order.Ship_via)
    t_purchase_order.shipping_method_id=<-shipping_method_id_chan


    /////////////////////////////////
	destination_country_id_chan :=make(chan string)
    go get_country_id_chan(destination_country_id_chan,t.Data.Purchase_order.Import_country)
    t_purchase_order.destination_country_id=<-destination_country_id_chan
////////////////////////////////////////
	// t_purchase_order.destination_country_id=t.Data.Purchase_order.Destination_country
	t_purchase_order.loading_port=t.Data.Purchase_order.Loading_port
	t_purchase_order.unloading_port=t.Data.Purchase_order.Unloading_port
	// t_purchase_order.certificate=t.Data.Purchase_order.Certificate
	t_purchase_order.po_url=t.Data.Purchase_order.Po_url
	t_purchase_order.total_quantity=t.Data.Purchase_order.Total_quantity
	t_purchase_order.total_amount=t.Data.Purchase_order.Total_amount
	// t_purchase_order.currency_id=t.Data.Purchase_order.Currency
	// ()
	currency_id_chan :=make(chan string)
    go get_currency_id(currency_id_chan,t.Data.Purchase_order.Currency)
    t_purchase_order.currency_id=<-currency_id_chan
    if t_purchase_order.currency_id==""{
    	return error_purchase_order_currency,errors.New("purchase_order currency_id missed")
    }
////////////////////////////////////////////////////////
	fmt.Println(t.Data.Purchase_order.Total_amount)
	fmt.Println(t.Data.Purchase_order.Currency)
	t_purchase_order.comments=t.Data.Purchase_order.Comments
	t_purchase_order.note=t.Data.Purchase_order.Note
	t_purchase_order.createAt=time.Now().Add(sd.company_time_zone).Format("2006-01-02 15:04:05")
	// fmt.Println(t_purchase_order.createAt)
	// t_purchase_order.createBy="go_fcgi"
	fmt.Println(t.Data.Purchase_order.Created_by)
	// t_purchase_order.createBy=t.Data.Purchase_order.Created_by
	
	system_account_id_chan :=make(chan string)
    go get_system_account_id_chan(system_account_id_chan,t.Data.Purchase_order.Created_by)
    t_purchase_order.createBy=<-system_account_id_chan
    
  	t_purchase_order.dr=0
  	t_purchase_order.data_version=1
  	return check_to_db(&t_purchase_order,t,sd)
}

func get_check_response(t *PoData) (string){
	received_chan :=make(chan string)
    go get_contact_account_id_sh_chan(received_chan,t.Data.Purchase_order.Supplier)
    received:=<-received_chan

	var sd=shared_data{}
	err_no,check_err:=po_check_data(t)
	if check_err!=nil{
		return `{"error_code":`+err_no+`,"error_msg":"`+check_err.Error()+`","data":{"bill_no":"`+t.Data.Purchase_order.Po_no+`","bill_type":"Purchase Order","receive_by":"`+received+`",   "company":"`+t.Data.Purchase_order.Company+`","receive_at":"`+t.Request_time+`"},"reply_time":"`+time.Now().Format("2006-01-02 15:04:05")+`"}`
	}
	s,err:=check_with_database(t,&sd,received)
	if err!=nil{
		return `{"error_code":`+s+`,"error_msg":"`+err.Error()+`","data":{"bill_no":"`+t.Data.Purchase_order.Po_no+`","bill_type":"Purchase Order","receive_by":"`+received+`",   "company":"`+t.Data.Purchase_order.Company+`","receive_at":"`+t.Request_time+`"},"reply_time":"`+time.Now().Format("2006-01-02 15:04:05")+`"}`
	}
	
	json_ret:=&Response_json{
		Error_code:200,Error_msg:"po received successfully at "+time.Now().Format("2006-01-02 15:04:05"),Data:Response_json_data{Bill_no:t.Data.Purchase_order.Po_no,Bill_type:t.Data.Purchase_order.Bill_type,Receive_by:received,Company:t.Data.Purchase_order.Company,Receive_at:t.Request_time},Reply_time:time.Now().Format("2006-01-02 15:04:05")}
	var buffer bytes.Buffer
    enc := json.NewEncoder(&buffer)

    err_encode := enc.Encode(json_ret)
    if err_encode!=nil{
    	return `{"error_code":`+error_json_encode+`,"error_msg":"`+err_encode.Error()+`","data":{"bill_no":"`+t.Data.Purchase_order.Po_no+`","bill_type":"Purchase Order","receive_by":"`+received+`",   "company":"`+t.Data.Purchase_order.Company+`","receive_at":"`+t.Request_time+`"},"reply_time":"`+time.Now().Format("2006-01-02 15:04:05")+`"}`
    }
	return buffer.String()
}