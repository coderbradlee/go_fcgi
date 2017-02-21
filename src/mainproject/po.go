 package main
 import (
    "log"
    "fmt"
    "encoding/json"
    "net/http"
    _"io/ioutil"
    "bytes"
    "time"
)
 const(
	error_json_decode="-100"
	error_json_encode="-101"
	error_db_insert="-102"
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
	        ret=`{"error_code":"`+error_json_decode+`","error_msg":`+err_decode.Error()+`,"data":{},"reply_time":"`+time.Now().Format("2006-01-02 15:04:05")+`"}`
	        fmt.Fprint(w,ret )
	        log.Printf("Started %s %s for %s:%s\nrespose:%s", r.Method, r.URL.Path, addr,"sbody",ret)
	        return;
	    }
	    // log.Println(t.Operation)
	    // var err_encode error
	    ret =get_response(&t)
	    // if err_encode != nil {
	    // 	// ret=`{"error_code":`+error_json_encode+`,"error_msg":`+err_encode.Error()+`,"data":{},"reply_time":"`+time.Now().Format("2006-01-02 15:04:05")+`"}`
	    //     // fmt.Fprint(w, ret)
	    //     // fmt.Println(ret)
	    //     // log.Fatal(err_encode.Error)
	    //     fmt.Fprint(w,ret )
	    //     log.Printf("Started %s %s for %s:%s\nrespose:%s\nerror:%s", r.Method, r.URL.Path, addr,"sbody",ret,err_encode.Error)
	    //     return;
	    // }
	    fmt.Fprint(w,ret )
	    log.Printf("Started %s %s for %s:%s\nrespose:%s", r.Method, r.URL.Path, addr,"sbody",ret)
	}

} 
func deal_with_database(t *DeliverGoodsForPO)error {
	var t_purchase_order purchase_order
		
	t_purchase_order.purchase_order_id=rand_string(20)
	t_purchase_order.po_no=t.Data.Purchase_order.Po_no
	t_purchase_order.po_date=t.Data.Purchase_order.Po_date
	t_purchase_order.status=t.Data.Purchase_order.Status

	//from t.Data.Purchase_order.Company find company_id
	t_purchase_order.company_id=get_company_id(t.Data.Purchase_order.Company)
	
	//from item_no find basic_id
	t_purchase_order.vendor_basic_id="vendor_basic_id"
	
	//待确定
	t_purchase_order.contact_account_id="contact_account_id"
	t_purchase_order.payment_terms=t.Data.Purchase_order.Payment_terms
	t_purchase_order.requested_delivery_date=t.Data.Purchase_order.Requested_delivery_date
	// t.Data.Purchase_order.Ship_via select id
	t_purchase_order.shipping_method_id=get_shipping_method_id(t.Data.Purchase_order.Ship_via)
	t_purchase_order.destination_country_id=t.Data.Purchase_order.Destination_country
	t_purchase_order.loading_port=t.Data.Purchase_order.Loading_port
	t_purchase_order.certificate=t.Data.Purchase_order.Certificate
	t_purchase_order.po_url=t.Data.Purchase_order.Po_url
	t_purchase_order.total_quantity=t.Data.Purchase_order.Total_quantity
	t_purchase_order.total_amount=t.Data.Purchase_order.Total_amount
	t_purchase_order.currency_id=t.Data.Purchase_order.Currency
	fmt.Println(t.Data.Purchase_order.Total_amount)
	fmt.Println(t.Data.Purchase_order.Currency)
	t_purchase_order.comments=t.Data.Purchase_order.Comments
	t_purchase_order.note=t.Data.Purchase_order.Note
	t_purchase_order.createAt=time.Now().Format("2006-01-02 15:04:05")
	// fmt.Println(t_purchase_order.createAt)
	t_purchase_order.createBy="go_fcgi"
  	t_purchase_order.dr=0
  	t_purchase_order.data_version=1
  	return insert_to_db(&t_purchase_order,t)
}
func insert_to_db(t_purchase_order* purchase_order,t *DeliverGoodsForPO)error {
    _, err := db.Exec(
        `INSERT INTO t_purchase_order(
	    purchase_order_id,po_no,po_date,status,company_id,vendor_basic_id,
		contact_account_id,payment_terms,requested_delivery_date,
		shipping_method_id,destination_country_id,loading_port,
		certificate,po_url,total_quantity,total_amount,currency_id,comments,
		note,createAt,createBy,dr,data_version) 
		VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,t_purchase_order.purchase_order_id,
			t_purchase_order.po_no,
			t_purchase_order.po_date,
			t_purchase_order.status,
			t_purchase_order.company_id,
			t_purchase_order.vendor_basic_id,
			t_purchase_order.contact_account_id,
			t_purchase_order.payment_terms,
			t_purchase_order.requested_delivery_date,
			t_purchase_order.shipping_method_id,
			t_purchase_order.destination_country_id,
			t_purchase_order.loading_port,
			t_purchase_order.certificate,
			t_purchase_order.po_url,
			t_purchase_order.total_quantity,
			t_purchase_order.total_amount,
			t_purchase_order.currency_id,
			t_purchase_order.comments,
			t_purchase_order.note,
			t_purchase_order.createAt,
			t_purchase_order.createBy,
		  	t_purchase_order.dr,
		  	t_purchase_order.data_version)
    if err!=nil{
    	return err
    }else{
    	return insert_purchase_order_detail(t_purchase_order,t)
    }
   
}

func get_response(t *DeliverGoodsForPO) (string){
	err:=deal_with_database(t)
	if err!=nil{
		// return `{"Error_code":"-300","Error_msg":"insert failed","Data":"","Reply_time":"2017-03-17 12:00:00"}`,err
		return `{"error_code":"`+error_db_insert+`","error_msg":"`+err.Error()+`","data":{},"reply_time":"`+time.Now().Format("2006-01-02 15:04:05")+`"}`
	}
	json_ret:=&Response_json{Error_code:"200",Error_msg:"Goods received successfully at "+time.Now().Format("2006-01-02 15:04:05"),Data:Response_json_data{Goods_receipt_no:t.Data.Purchase_order.Po_no,Bill_type:t.Data.Purchase_order.Bill_type,Receive_by:"ERP",Company:t.Data.Purchase_order.Company,Receive_at:time.Now().Format("2006-01-02 15:04:05")},Reply_time:time.Now().Format("2006-01-02 15:04:05")}
		
	var buffer bytes.Buffer
    enc := json.NewEncoder(&buffer)

    err_encode := enc.Encode(json_ret)
    if err_encode!=nil{
    	return `{"error_code":"`+error_json_encode+`","error_msg":"`+err_encode.Error()+`","data":{},"reply_time":"`+time.Now().Format("2006-01-02 15:04:05")+`"}`
    }
	return buffer.String()
}