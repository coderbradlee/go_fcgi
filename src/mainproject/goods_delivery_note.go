 package main
 import (
)
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
func insert_goods_delivery_note(t *purchase_order,origi *DeliverGoodsForPO)error {
    var err error
    for _,deliver_notes:= range origi.Data.Deliver_notes{
        _, err = db.Exec(
        `INSERT INTO t_goods_delivery_note(
        note_id,goods_delivery_note_no,bill_type_id,company_id,
        purchase_order_id,buyer_id,vendor_master_id,fulfill_status,
        destination_port,trade_term_id,transport_term_id,packing_method_id,
        logistic_provider_master_id,logistic_provider_contact_id,etd,
        eta,atd,ata,customs_clearance_date,receiver,total_freight_charges,
        total_insurance_fee,total_excluded_tax,note,createAt,createBy,dr,
        data_version) 
        VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
        rand_string(20),
        "goods_delivery_note_no",
        "bill_type_id",//get_bill_type_id(t.Bill_type)
        t.company_id,
        t.purchase_order_id,
        "buyer_id",//get_buyer_id(deliver_notes.buyer)
        "vendor_master_id",//get_vendor_master_id(t.vendor_basic_id)
        t.Status,//
        deliver_notes.Loading_port,
        "trade_term_id",//get_trade_term_id(deliver_notes.Trade_term)
        "transport_term_id",
        packing_method_id,//get_packing_method_id(deliver_notes.Packing_method)
        logistic_provider_master_id,//get_logistic_master_id(deliver_notes.Logistic)
        logistic_provider_contact_id,//get_logistic_contact_id(deliver_notes.Logistic_contact)
        deliver_notes.Etd,
        deliver_notes.Eta,
        "atd",
        "ata",
        deliver_notes.Customs_clearance_date,
        "receiver",
        deliver_notes.Total_freight_charges,
        deliver_notes.Total_insurance_fee,
        deliver_notes.Total_excluded_tax,
        "note",
        time.Now().Format("2006-01-02 15:04:05"),
        "go_fcgi",
        0,
        1)
    }
    return err
}