 package main
 import (
    "log"
    "fmt"
    "encoding/json"
    "net/http"
    _"io/ioutil"
    _"bytes"
    "os"
)
type Detail struct{
	product_name string
	product_code string
	item_no string
	unit_price float64
	quantity int32
	uom string
	sub_total float64
	warranty int32
	comments string
	note string
}
type Purchase_order struct{
	bill_type string
	po_no string
	po_url string
	po_date string
	create_by string
	status int32
	supplier string
	website string
	company string
	requested_delivery_date string
	trade_term  string
	payment_terms string
	ship_via string
	destination_country string
	loading_port string
	certificate string
	total_quantity int32
	total_amount float64
	currency string
	comments string
	note string
	detail []Detail
}
type Commercial_invoice struct{
	ci_no string
	ci_url string
}
type Packing_list struct{
	pl_no string
	pl_url string
}
type Bill_of_lading struct{
	bl_no string
	bl_url string
}
type Associated_so struct{
	associated_so_no string
	associated_so_url string
}

type Deliver_notes_detail struct{
	product_name string
	product_code string
	item_no string
	unit_price float64
	quantity int32
	uom string
	sub_total float64
}
type Deliver_notes struct{
	supplier string
	buyer string
	loading_port string
	trade_term string
	ship_via string
	packing_method string
	logistic string
	logistic_contact string
	logistic_contact_email string
	logistic_contact_telephone_number string
	etd string
	eta string
	customs_clearance_date string
	total_freight_charges float64
	total_insurance_fee float64
	total_excluded_tax float64
	currency string
	commercial_invoice Commercial_invoice
	packing_list Packing_list
	bill_of_lading Bill_of_lading
	associated_so Associated_so
	detail []Deliver_notes_detail

}
type Data struct{
	request_system int32
	request_time string
	purchase_order Purchase_order
	deliver_notes []Deliver_notes
}
type DeliverGoodsForPO struct {
   operation string
   data Data 
}
type Response_json_data struct{
	goods_receipt_no string
	bill_type string
	receive_by string
	company string
	receive_at string
}
type Response_json struct{
	error_code string
	error_msg string
	data Response_json_data	
	reply_time string		   
}
func poHandler (w http.ResponseWriter, r *http.Request) {
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
 		var ret=""
		// // log.Println(sbody)
		// log.Printf("Started %s %s for %s:%s", r.Method, r.URL.Path, addr,sbody)
		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()
	    var t DeliverGoodsForPO   
	    err_decode := decoder.Decode(&t)
	    if err_decode != nil {
	        // panic(err)
	        ret=`{"error_code":"-100","error_msg":"json decoder error","data":{},"reply_time":"2017-03-17 12:00:00"}`
	    }
	    log.Println(t.operation)

	    // response_data:=Response_json_data{goods_receipt_no:"GR-FR-20170226-000196",bill_type:"Goods Receipt",receive_by:"Enie Yang",company:"ReneSola France",receive_at:"2017-03-17 12:00:00"}
	    json_ret:=&Response_json{error_code:"200",error_msg:"Goods received successfully at 2017-03-17 12:00:00",data:Response_json_data{goods_receipt_no:"GR-FR-20170226-000196",bill_type:"Goods Receipt",receive_by:"Enie Yang",company:"ReneSola France",receive_at:"2017-03-17 12:00:00"},reply_time:"2017-03-17 12:00:00"}
		// fmt.Println(json_ret)
	    // encoder:=json.NewEncoder(w)
	    // err_encode:=encoder.Encode(json_ret)
	    // if err_encode != nil {
	    //     // panic(err)
	    //     ret=`{"error_code":"-200","error_msg":"json encoder error","data":{},"reply_time":"2017-03-17 12:00:00"}`
	    //     fmt.Fprint(w, ret)
	    // }
		// var buffer bytes.Buffer
	    // enc := json.NewEncoder(&buffer)

	    // err_encode := enc.Encode(json_ret)
	    buffer,err_encode:=json.Marshal(json_ret)
	    if err_encode != nil {
	    	ret=`{"error_code":"-200","error_msg":"json encoder error","data":{},"reply_time":"2017-03-17 12:00:00"}`
	        fmt.Fprint(w, ret)
	        fmt.Println("error encoding the response to a join request")
	        log.Fatal(err_encode.Error)
	    }
	    os.Stdout.Write(buffer)
		fmt.Printf("response: %s\n", buffer)
	    fmt.Fprint(w, buffer)
	    // fmt.Printf("response: %s\n", buffer.Bytes())
	    // fmt.Fprint(w, buffer.Bytes())
	}

} 
