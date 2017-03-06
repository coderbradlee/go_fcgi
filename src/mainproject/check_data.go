 package main
 import (
    // "time"
    "errors"
    "os"
    "fmt"
    // "reflect"
)
const(
    error_json_decode="-100"//json 解包错误
    error_json_encode="-101"//json 打包错误
    error_db="-102"//连接mysql数据库错误
    error_check_request_system="-120"//请求系统是否为1
    error_check_bill_type="-121"//bill_tpye是否为Purchase Order
    // error_check_po_no="-122"//po_no长度问题，数据库表里面是20位
    error_check_po_url="-123"//是否存在文件
    error_check_status="-124"//status是否为1
    error_check_supplier="-125"//supplier是否为Renesola Shanghai
    error_check_packing_method="-126"//t_packing_method表里没有此packing_method
    error_check_logistic_provider="-127"//物流提供商信息缺失
    //add more error
    error_check_ship_via="-128"
    error_check_trade_term="-129"
    error_check_payment_terms="-130"

    error_deliver_notes_packing_method_id="-140"
    error_deliver_notes_transport_term_id="-141"
    error_deliver_notes_buyer_id="-142"
    error_deliver_notes_trade_term_id="-143"
    error_deliver_notes_vendor_master_id="-144"
    // error_deliver_notes_bill_type_id="-145"

    error_purchase_order_detail_uom_id="-146"
    error_purchase_order_detail_item_master_id="-147"

    error_deliver_notes_logistic_master_id="-148"

    error_insert_purchase_order="-150"
    error_insert_purchase_order_detail="-151"
    error_insert_goods_delivery_note="-152"
    error_insert_goods_delivery_note_attachment="-153"
    error_insert_goods_delivery_note_detail="-154"
    error_insert_goods_receipt="-155"
    error_insert_commercial_invoice="-156"
)
type check_struct struct{
    error_code string
    err error
}
func check_request_system(request_system int32,error_chan chan<- check_struct) {
    var t check_struct
    if request_system!=1{
        t=check_struct{error_check_request_system,errors.New("request_system !=1")}
    }
    error_chan<- t
}
func check_bill_type(bill_type string,error_chan chan<- check_struct) {
    var t check_struct
    if bill_type!="Purchase Order"{
        t=check_struct{error_check_bill_type,errors.New("bill_type!=Purchase Order")}
    }
    error_chan<- t
}
// func check_po_no(po_no string,error_chan chan<- map[string]error) {
//     t:=make(map[string]error)
//     cs:="PO-FR-20170216-00101"
//     if len(po_no)>len(cs){
//         error_chan<- errors.New(`po_no is too long`)
//     }
//     error_chan<- t
// }
func check_po_url(po_url string,error_chan chan<- check_struct) {
    var t check_struct
    _,err:=os.Stat(po_url)
    t=check_struct{error_check_po_url,err}
    error_chan<- t
}
func check_status(status int32,error_chan chan<- check_struct) {
    var t check_struct
    if status!=1{
        t=check_struct{error_check_status,errors.New(`status!=1`)}
    }
    error_chan<- t
}
func check_supplier(supplier string,error_chan chan<- check_struct) {
    var t check_struct
    if supplier!="Renesola Shanghai"{
        t=check_struct{error_check_supplier,errors.New(`supplier is not Renesola Shanghai`)}
    }
    error_chan<- t
}
func check_packing_method(deliver_notes []Deliver_notes,error_chan chan<- check_struct) {
    var t check_struct
    for _,d:=range deliver_notes{
        // fmt.Println(reflect.TypeOf(d))
        var packing_method string
        db.QueryRow("select packing_method_id from t_packing_method where name=?",d.Packing_method).Scan(&packing_method)
        if packing_method== ""{
            t=check_struct{error_check_packing_method,errors.New(`packing_method_id missed`)}
        }
    }
    error_chan<- t
}
func check_logistic_provider(deliver_notes []Deliver_notes,error_chan chan<- check_struct) {
    var t check_struct
    for _,d:=range deliver_notes{
        // fmt.Println(reflect.TypeOf(d))
        var logistic_provider_basic_id string
        db.QueryRow("select logistic_provider_basic_id from t_logistic_provider_basic where name=?",d.Logistic).Scan(&logistic_provider_basic_id)
        if logistic_provider_basic_id== ""{
            t=check_struct{error_check_logistic_provider,errors.New(`logistic_provider_basic_id missed`)}
        }
    }
    error_chan<- t
}
// func check_ship_via(ship_via string,error_chan chan<- check_struct) {
//     var transport_term_id string
//     db.QueryRow("select ship_via_id from t_ship_via where full_name=?",ship_via).Scan(&transport_term_id)
//     transport_term_id_chan<-transport_term_id
// }
func check_ship_via(ship_via string,error_chan chan<- check_struct) {
    var t check_struct
    var transport_term_id string
    db.QueryRow("select ship_via_id from t_ship_via where full_name=?",ship_via).Scan(&transport_term_id)
    if transport_term_id== ""{
        t=check_struct{error_check_ship_via,errors.New(`ship_via missed`)}
    }
    error_chan<- t
}
func check_trade_term(Trade_term string,error_chan chan<- check_struct) {
    var t check_struct
    var trade_term_id string
    db.QueryRow("select trade_term_id from t_trade_term where short_name=?",Trade_term).Scan(&trade_term_id)
    if trade_term_id== ""{
        t=check_struct{error_check_trade_term,errors.New(`trade_term_id missed`)}
    }
    error_chan<- t
}

func check_data(origi *DeliverGoodsForPO)(string,error) {
    // var all_error map[string]error
    error_chan:=make(chan check_struct)
    go check_request_system(origi.Data.Request_system,error_chan)
    go check_bill_type(origi.Data.Purchase_order.Bill_type,error_chan)
    go check_po_url(origi.Data.Purchase_order.Po_url,error_chan)
    go check_status(origi.Data.Purchase_order.Status,error_chan)
    go check_supplier(origi.Data.Purchase_order.Supplier,error_chan)
    go check_packing_method(origi.Data.Deliver_notes,error_chan)
    go check_logistic_provider(origi.Data.Deliver_notes,error_chan)
    go check_ship_via(origi.Data.Purchase_order.Ship_via,error_chan)
    go check_trade_term(origi.Data.Purchase_order.Trade_term,error_chan)
    for i:=0;i<9;i++{
        err:=<-error_chan
        fmt.Println("104:",err.error_code,err.err)
        if err.err!=nil{
            return err.error_code,err.err
        }
    } 
    return "",nil
}
