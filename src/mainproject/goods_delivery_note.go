 package main
 import (
    "time"
    "log"
    "fmt"
    "net/http"
    "io/ioutil"
    "strconv"
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
func get_bill_type_id()string {
    var bill_type_id string
    db.QueryRow("select bill_type_id from t_bill_type where code='GDN'").Scan(&bill_type_id)
    return bill_type_id
}
func get_vendor_master_id(vendor_basic_id string)string {
    var vendor_master_id string
    db.QueryRow("select vendor_master_id from t_vendor_master where vendor_basic_id=?",vendor_basic_id).Scan(&vendor_master_id)
    return vendor_master_id
}
func get_trade_term_id(Trade_term string)string {
    var trade_term_id string
    db.QueryRow("select trade_term_id from t_trade_term where short_name=?",Trade_term).Scan(&trade_term_id)
    return trade_term_id
}
func get_transport_term_id(ship_via string)string {
    var transport_term_id string
    db.QueryRow("select ship_via_id from t_ship_via where full_name=?",ship_via).Scan(&transport_term_id)
    return transport_term_id
}
func get_packing_method_id(Packing_method string)string {
    var packing_method_id string
    db.QueryRow("select packing_method_id from t_packing_method where name=?",Packing_method).Scan(&packing_method_id)
    return packing_method_id
}
func get_logistic_master_id(Logistic string)string {
    var logistic_master_id string
    db.QueryRow("select logistic_provider_master_id from t_logistic_provider_master where native_name=?",Logistic).Scan(&logistic_master_id)
    return logistic_master_id
}
func get_logistic_contact_id(Logistic_contact string)string {
    var logistic_contact_id string
    //db.QueryRow("select logistic_contact_id from t_logistic_provider_master where native_name=?",Logistic_contact).Scan(&logistic_contact_id)
    return logistic_contact_id
}
func get_buyer_id(buyer string)string {
    var buyer_id string
    //db.QueryRow("select logistic_contact_id from t_logistic_provider_master where native_name=?",Logistic_contact).Scan(&logistic_contact_id)
    return buyer_id
}
type flow_no_json struct{
    FlowNo string `json:"flowNo"`
    ReplyTime string `json:"replyTime"`
}
func get_flow_no(company string)(string,error) {
    // var flow_no string
    //http://127.0.0.1:8088/flowNo/JP/SO
    url:=configuration.Redis_url+"/"+company+"/PO"

    resp, err1 := http.Get(url)
    if err1 != nil {
        return  "",err1
    }

    defer resp.Body.Close()
    body, err2 := ioutil.ReadAll(resp.Body)
    if err2 != nil {
        // handle error
        return  "",err2
    }
    var data flow_no_json
    json.Unmarshal(body, &data)
    i, err3 := strconv.Atoi(data.FlowNo))
    if err3 != nil {
        // handle error
        return  "",err3
    }
    // str := string.format(%06d",i)
    str := fmt.Sprintf("%06d",i)
    return str,nil
}
func get_goods_delivery_note_no(company string)(string,error) {
    goods_delivery_note_no:="PO-"
    var short string
    db.QueryRow("select note from t_company where short_name=?",company).Scan(&short)
    // QU-UK-20160930-000001
    goods_delivery_note_no+=short+"-"
    goods_delivery_note_no+=time.Now().Format("20060102")+"-"
    flow,err:=get_flow_no(short)
    if err!=nil{
        return "",err
    }
    goods_delivery_note_no+=flow//get int,format to 6bit,then convert to string
    return goods_delivery_note_no,nil
}
func insert_goods_delivery_note(t *purchase_order,origi *DeliverGoodsForPO)error {
    var err error
    // log.Println("insert_goods_delivery_note")
    for _,deliver_notes:= range origi.Data.Deliver_notes{
        // bill_type_id:=get_bill_type_id(t.Bill_type)
        bill_type_id:=get_bill_type_id()
        buyer_id:=get_buyer_id(deliver_notes.Buyer)
        vendor_master_id:=get_vendor_master_id(t.vendor_basic_id)
        trade_term_id:=get_trade_term_id(deliver_notes.Trade_term)
        transport_term_id:=get_transport_term_id(deliver_notes.Ship_via)
        packing_method_id:=get_packing_method_id(deliver_notes.Packing_method)
        logistic_master_id:=get_logistic_master_id(deliver_notes.Logistic)
        logistic_contact_id:=get_logistic_contact_id(deliver_notes.Logistic_contact)
        goods_delivery_note_no,err:=get_goods_delivery_note_no(origi.Data.Purchase_order.Company)
        if err!=nil{
            return err
        }
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
        goods_delivery_note_no,//goods_delivery_note_no 待定
        bill_type_id,
        t.company_id,
        t.purchase_order_id,
        buyer_id,
        vendor_master_id,
        t.status,
        deliver_notes.Loading_port,
        trade_term_id,
        transport_term_id,//transport_term_id 待定
        packing_method_id,
        logistic_master_id,
        logistic_contact_id,
        deliver_notes.Etd,
        deliver_notes.Eta,
        "",//atd
        "",//ata
        deliver_notes.Customs_clearance_date,
        "",//receiver 待定
        deliver_notes.Total_freight_charges,
        deliver_notes.Total_insurance_fee,
        deliver_notes.Total_excluded_tax,
        "",//note
        time.Now().Format("2006-01-02 15:04:05"),
        "go_fcgi",
        0,
        1)
    }
    if err!=nil{
        log.Println("insert_goods_delivery_note:", err) 
    }
    return err
}