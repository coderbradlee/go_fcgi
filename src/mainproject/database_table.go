 package main
func rand_string(lens int)string{
    choice:="ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789";
    var ret string
    rand.Seed(time.Now().UnixNano())
    for i:=0;i<lens;i++{
        ret+=string(choice[rand.Intn(35)]);
    }
   return ret;
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