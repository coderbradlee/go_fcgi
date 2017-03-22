package main
import (
    // "math/rand"
    // "time"
)

/**
 * json data define
 */
type Commercial_invoice struct{
	Ci_no string `json:"ci_no"`
	Ci_url string `json:"ci_url"`
	Ci_date string `json:"ci_date"`
	Status int32 `json:"status"`
	Company string `json:"company"`
	Invoice_type int32 `json:"invoice_type"`
	Total_amount float64 `json:"total_amount"`
	Currency string `json:"currency"`
	Created_by string `json:"created_by"`
	Approved_by string `json:"approved_by"`
	Note string `json:"note"`
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
	Company string `json:"company"`
	Bill_type string `json:"bill_type"`
	Gdn_no string `json:"gdn_no"`
	Po_no string `json:"po_no"`
	Supplier string `json:"supplier"`
	Buyer string `json:"buyer"`
	Trade_term string `json:"trade_term"`
	Ship_via string `json:"ship_via"`
	Packing_method string `json:"packing_method"`
	Export_country string `json:"export_country"`
	Loading_port string `json:"loading_port"`
	Import_country string `json:"import_country"`
	Unloading_port string `json:"unloading_port"`

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
	Created_by string `json:"created_by"`
    Approved_by string `json:"approved_by"`
	Comments string `json:"comments"`
	Note string `json:"note"`
}
type DeliverGoodsData struct{	
	Deliver_notes []Deliver_notes `json:"deliver_notes"`
}
type DeliverGoodsForPO struct {
	Request_system int32 `json:"request_system"`
	Request_time string `json:"request_time"`
   	Operation string `json:"operation"`
   	Data DeliverGoodsData  `json:"data"`
}