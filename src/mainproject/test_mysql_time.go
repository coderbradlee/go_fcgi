 package main
 import (
    "logger"
    "fmt"
    "encoding/json"
    "net/http"
    "io/ioutil"
    // "bytes"
    // "time"
    // "errors"
    // "runtime/pprof"
)
func test_mysql_time (w http.ResponseWriter, r *http.Request) {
	////////////////////////////////
	addr := r.Header.Get("X-Real-IP")
	if addr == "" {
		addr = r.Header.Get("X-Forwarded-For")
		if addr == "" {
			addr = r.RemoteAddr
		}
	}
	if r.Method !="POST"{
		fmt.Fprint(w, "this interface should be post!")
	} else{
		body, _:= ioutil.ReadAll(r.Body)
	    
	    // log.Println(string(body))
	    var t PoData  
	    // var po_shared_data shared_data
	    json.Unmarshal(body, &t)
	    // bytes.Trim(body,"\\r\\n")
	    // line := strings.Trim(string(body), "\r\n")
		defer r.Body.Close()
		insert_test_time()
		ret:=single_select2()+"error_codeerror_codeerror_codeerror_codeerror_codeerror_codeerror_codeerror_codeerror_codeerror_codeerror_codeerror_codeerror_codeerror_codeerroerror_codeerror_codeerror_codeerror_codeerror_codeerror_codeerror_codeerror_codeerror_codeerror_codeerror_codeerror_codeerror_codeerror_codeerror_codeerror_code"
	    fmt.Fprint(w,ret )
	    log_str:=fmt.Sprintf("Started %s %s for %s:%s response:%s", r.Method, r.URL.Path, addr,"body",ret)
        logger.Info(log_str)
	}

} 

func single_select()string {
	var packing_method string
	for i:=0;i<10;i++{
		db.QueryRow("select packing_method_id from t_packing_method where name=?","Pallet").Scan(&packing_method)
	}
    return packing_method
}
func single_select2()string {
	var purchase_order_id string
	for i:=0;i<10;i++{
		purchase_order_id_chan :=make(chan string)
        go get_purchase_order_id_chan(purchase_order_id_chan,"PO-FR-20170216-0016")
        purchase_order_id=<-purchase_order_id_chan
	}
    return purchase_order_id
}
func insert_test_time() {
	db.Exec(
        `INSERT INTO t_purchase_order(
	    purchase_order_id,po_no,po_date,status,company_id,vendor_basic_id,
		contact_account_id,payment_terms,requested_delivery_date,
		shipping_method_id,destination_country_id,loading_port,unloading_port,
		certificate,po_url,total_quantity,total_amount,currency_id,comments,
		note,createAt,createBy,updateBy,dr,data_version) 
		VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,"purchase_order_id",
			"po_no",
			"po_date",
			5,// t_purchase_order.status,
			"company_id",
			"vendor_basic_id",
			"contact_account_id",
			"payment_terms",
			"requested_delivery_date",
			"shipping_method_id",
			"destination_country_id",
			"loading_port",
			"unloading_port",
			"certificate",
			"po_url",
			"total_quantity",
			"total_amount",
			"currency_id",
			"comments",
			"note",
			"createAt",
			"createBy",
			"go_fcgi",
		  	0,
		  	1)
	db.Exec("delete from t_purchase_order where purchase_order_id='purchase_order_id'")    
}