package main
import (
    "math/rand"
    "time"
)
func rand_string(lens int)string{
    choice:="ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789";
    var ret string
    rand.Seed(time.Now().UnixNano())
    for i:=0;i<lens;i++{
        ret+=string(choice[rand.Intn(35)]);
    }
   return ret;
}
type shared_data struct{
	goods_receipt_no string//在response中回传发货号
	goods_delivery_note_id string//在插入t_goods_delivery_note时生成，在插入t_goods_delivery_note_detail和t_goods_delivery_note_attachment的时候用到
	company_time_zone time.Duration //获取后用于在各个表中的createAt字段
	                  
}
type purchase_order struct{
	purchase_order_id string/*主键*/
	po_no string/*采购订单编号*/
	po_date string/*下单日期，YYYY-MM-DD*/
	status int32/*状态 CRM使用：0:Canceled(已取消); 1:UnConfirmed(未确认); ERP使用: 4:Confirmed(已确认); 5:Refused(已拒绝); 6:Inbounded(已入库); 7:Delivered(已交付); 8:Part_Delivered(部分交付)*/
	company_id string/*采购公司id*/
	vendor_basic_id string/*供应商基本档案id*/
	contact_account_id string/*联系人帐号id*/
	payment_terms string/*付款条款*/
	requested_delivery_date string/*要求发货日期，YYYY-MM-DD*/
	shipping_method_id string/*运输方式id*/
	destination_country_id string/*目的地国家id*/
	loading_port string/*到货港口*/
	certificate string/*证书*/
	po_url string/*采购订单PDF文件地址*/
	total_quantity int32/*订单总数量*/
	total_amount float64/*订单总金额*/
	currency_id string/*交易币种id*/
	comments string/*说明*/
	note string/*备注*/
	createAt string/*创建时间, YYYY-MM-DD HH:MM:SS*/
	createBy string/*创建人*/
  	dr int32/*逻辑删除标志 0:未删除;1:已删除*/
  	data_version int32/*版本号, 从0开始*/
}

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
	Bill_type string `json:"bill_type"`
	Po_no string `json:"po_no"`
	Po_url string `json:"po_url"`
	Po_date string `json:"po_date"`
	Created_by string `json:"created_by"`
	Approved_by string `json:"approved_by"`
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
	Currency string `json:"currency"`
	Comments string `json:"comments"`
	Note string `json:"note"`
	Detail []Detail `json:"detail"`
}
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
	Deliver_note_no string `json:"deliver_note_no"`
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
	Reply_system int32 `json:"reply_system"`
}
type Response_json struct{
	Error_code string `json:"error_code"`
	Error_msg string `json:"error_msg"`
	Data Response_json_data	 `json:"response_json_data"`
	Reply_time string `json:"reply_time"`		   
}
type Delivery_note struct{
    note_id string/*主键*/
    goods_delivery_note_no string/*发货单号*/
    bill_type_id string/*单据类型id*/
    company_id string/*公司id*/
    purchase_order_id string/*采购订单id*/
    buyer_id string/*采购员id*/
    vendor_master_id string/*供应商管理档案id*/
    fulfill_status int32/*履行状态, 0:pending(待处理); 1:scheduled(已调度); 2:received(已收货);*/
    destination_port string/*目的地港口*/
    trade_term_id string/*贸易方式id*/
    transport_term_id string/*运输方式id*/
    packing_method_id string/*包装方式id*/
    logistic_provider_master_id string/*物流商id*/
    logistic_provider_contact_id string/*物流商联系人id*/
    etd string/*预计离港时间*/
    eta string/*预计到港时间*/
    atd string/*实际离港时间*/
    ata string/*实际到港时间*/
    customs_clearance_date string/*海关清关时间*/
    receiver string/*通知接收人id*/
    total_freight_charges float64/*总运费金额*/
    total_insurance_fee float64/*总保险费*/
    total_excluded_tax float64/*商品总金额（不含税）*/
    note string/*备注*/
    createAt string
    createBy string
    dr int32
    data_version int32
}