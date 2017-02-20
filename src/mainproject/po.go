 package main
 import (
    "log"
    "fmt"
    "encoding/json"
    "net/http"
    _"io/ioutil"
    "bytes"
)
type Detail struct{
	Product_name string `json:"product_name"`
	Product_code string `json:"product_code"`
	Item_no string `json:"item_no"`
	Unit_price float64 `json:"unit_price"`
	Quantity int32 `json:"quantity"`
	Uom string `json:"uom"`
	Sub_total float64 `json:"sub_total"`
	Warranty int32 `json:"warranty"`
	Comments string `json:"comments"`
	Note string `json:"note"`
}
type Purchase_order struct{
	Bill_type string `json:"bill_type"`
	Po_no string `json:"po_no"`
	Po_url string `json:"po_url"`
	Po_date string `json:"po_date"`
	Create_by string `json:"create_by"`
	Status int32 `json:"status"`
	Supplier string `json:"supplier"`
	Website string `json:"website"`
	Company string `json:"company"`
	Requested_delivery_date string `json:"requested_delivery_date"`
	Trade_term  string `json:"trade_term"`
	Payment_terms string `json:"payment_terms"`
	Ship_via string `json:"ship_via"`
	Destination_country string `json:"destination_country"`
	Loading_port string `json:"loading_port"`
	Certificate string `json:"certificate"`
	Total_quantity int32 `json:"total_quantity"`
	Total_amount float64 `json:"total_amount"`
	Currency string `json:"total_amount"`
	Comments string `json:"comments"`
	Note string `json:"note"`
	Detail []Detail `json:"detail"`
}
type Commercial_invoice struct{
	Ci_no string `json:"ci_no"`
	Ci_url string `json:"ci_url"`
}
type Packing_list struct{
	Pl_no string `json:"pl_no"`
	Pl_url string `json:"pl_url"`
}
type Bill_of_lading struct{
	Bl_no string `json:"bl_no"`
	Bl_url string `json:"bl_url"`
}
type Associated_so struct{
	Associated_so_no string `json:"associated_so_no"`
	Associated_so_url string `json:"associated_so_url"`
}

type Deliver_notes_detail struct{
	Product_name string `json:"product_name"`
	Product_code string `json:"product_code"`
	Item_no string `json:"item_no"`
	Unit_price float64 `json:"unit_price"`
	Quantity int32 `json:"quantity"`
	Uom string `json:"uom"`
	Sub_total float64 `json:"sub_total"`
}
type Deliver_notes struct{
	Supplier string `json:"supplier"`
	Buyer string `json:"buyer"`
	Loading_port string `json:"loading_port"`
	Trade_term string `json:"trade_term"`
	Ship_via string `json:"ship_via"`
	Packing_method string `json:"packing_method"`
	Logistic string `json:"logistic"`
	Logistic_contact string `json:"logistic_contact"`
	Logistic_contact_email string `json:"logistic_contact_email"`
	Logistic_contact_telephone_number string `json:"logistic_contact_telephone_number"`
	Etd string `json:"etd"`
	Eta string `json:"eta"`
	Customs_clearance_date string `json:"customs_clearance_date"`
	Total_freight_charges float64 `json:"total_freight_charges"`
	Total_insurance_fee float64 `json:"total_insurance_fee"`
	Total_excluded_tax float64 `json:"total_excluded_tax"`
	Currency string `json:"currency"`
	Commercial_invoice Commercial_invoice `json:"commercial_invoice"`
	Packing_list Packing_list `json:"packing_list"`
	Bill_of_lading Bill_of_lading `json:"bill_of_lading"`
	Associated_so Associated_so `json:"associated_so"`
	Detail []Deliver_notes_detail `json:"detail"`

}
type Datas struct{
	Request_system int32 `json:"request_system"`
	Request_time string `json:"request_time"`
	Purchase_order Purchase_order `json:"purchase_order"`
	Deliver_notes []Deliver_notes `json:"deliver_notes"`
}
type DeliverGoodsForPO struct {
   Operation string `json:"operation"`
   Data Datas  `json:"data"`
}
type Response_json_data struct{
	Goods_receipt_no string `json:"goods_receipt_no"`
	Bill_type string `json:"bill_type"`
	Receive_by string `json:"receive_by"`
	Company string `json:"company"`
	Receive_at string `json:"receive_at"`
}
type Response_json struct{
	Error_code string `json:"error_code"`
	Error_msg string `json:"error_msg"`
	Data Response_json_data	 `json:"response_json_data"`
	Reply_time string `json:"reply_time"`		   
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
	        fmt.Fprint(w,ret )
	        log.Printf("Started %s %s for %s:%s\nrespose:%s", r.Method, r.URL.Path, addr,"sbody",ret)
	        return;
	    }
	    log.Println(t.Operation)

	    
	    ret,err_encode=get_response()
	    if err_encode != nil {
	    	ret=`{"error_code":"-200","error_msg":"json encoder error","data":{},"reply_time":"2017-03-17 12:00:00"}`
	        // fmt.Fprint(w, ret)
	        // fmt.Println(ret)
	        // log.Fatal(err_encode.Error)
	        fmt.Fprint(w,ret )
	        log.Printf("Started %s %s for %s:%s\nrespose:%s", r.Method, r.URL.Path, addr,"sbody",ret)
	        return;
	    }
	    fmt.Fprint(w,ret )
	    log.Printf("Started %s %s for %s:%s\nrespose:%s", r.Method, r.URL.Path, addr,"sbody",ret)
	}

} 
func get_response() (string, error){
	json_ret:=&Response_json{Error_code:"200",Error_msg:"Goods received successfully at 2017-03-17 12:00:00",Data:Response_json_data{Goods_receipt_no:"GR-FR-20170226-000196",Bill_type:"Goods Receipt",Receive_by:"Enie Yang",Company:"ReneSola France",Receive_at:"2017-03-17 12:00:00"},Reply_time:"2017-03-17 12:00:00"}
		
		var buffer bytes.Buffer
	    enc := json.NewEncoder(&buffer)

	    err_encode := enc.Encode(json_ret)
	return buffer.String(),err_encode
}