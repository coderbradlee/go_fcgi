package main
import (
    "math/rand"
    "time"
)

/**
 * json data define
 */
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
	Website string `json:"website"`
	Company string `json:"company"`
	Bill_type string `json:"bill_type"`
	Po_no string `json:"po_no"`
	Po_url string `json:"po_url"`
	Po_date string `json:"po_date"`
	Created_by string `json:"created_by"`
	Approved_by string `json:"approved_by"`
	Status int32 `json:"status"`
	Supplier string `json:"supplier"`
	
	Requested_delivery_date string `json:"requested_delivery_date"`
	Trade_term  string `json:"trade_term"`
	Payment_terms string `json:"payment_terms"`
	Ship_via string `json:"ship_via"`

	Export_country string `json:"export_country"`
	Loading_port string `json:"loading_port"`
	Import_country string `json:"import_country"`
	Unloading_port string `json:"unloading_port"`

	Certificate string `json:"certificate"`
	Total_quantity int32 `json:"total_quantity"`
	Total_amount float64 `json:"total_amount"`
	Currency string `json:"currency"`
	Comments string `json:"comments"`
	Note string `json:"note"`
	Detail []Detail `json:"detail"`
}

type Datas struct{
	Request_system int32 `json:"request_system"`
	Request_time string `json:"request_time"`
	Purchase_order Purchase_order `json:"purchase_order"`
}
type PoData struct {
   Operation string `json:"operation"`
   Data Datas  `json:"data"`
}
type Response_json_data struct{
	Bill_no string `json:"bill_no"`
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
