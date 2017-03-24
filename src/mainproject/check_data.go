 package main
 import (
    // "time"
    "errors"
    "os"
    // "fmt"
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
    error_check_po_exists="-126"//t_packing_method表里没有此packing_method
    error_check_logistic_provider="-127"//物流提供商信息缺失
    //add more error
    error_check_ship_via="-128"
    error_check_trade_term="-129"
    error_check_payment_terms="-130"

    error_check_deliver_notes_commercial_invoice="-131"
    error_check_deliver_notes_packing_list="-132"
    error_check_deliver_notes_bill_of_lading="-133"
    error_check_deliver_notes_associated_so="-134"
    error_check_deliver_notes_deliver_note_no="-135"
    error_check_po_exist_for_gdn="-136"


    error_deliver_notes_packing_method_id="-140"
    error_deliver_notes_transport_term_id="-141"
    error_deliver_notes_buyer_id="-142"
    error_deliver_notes_trade_term_id="-143"
    error_deliver_notes_vendor_master_id="-144"
    
    error_purchase_order_currency="-145"
    error_purchase_order_detail_uom_id="-146"
    error_purchase_order_detail_item_master_id="-147"

    // error_deliver_notes_logistic_master_id="-148"
    error_commercial_invoice_status="-148"

    //以下为插入数据库表时报错
    error_insert_purchase_order="-150"
    error_insert_purchase_order_detail="-151"
    error_insert_goods_delivery_note="-152"
    error_insert_goods_delivery_note_attachment="-153"
    error_insert_goods_delivery_note_detail="-154"
    error_insert_goods_receipt="-155"
    error_insert_commercial_invoice="-156"
    error_get_flow_no_po="-157"
    error_get_flow_no_gdn="-158"
    error_call_erp_api="-159"
)
type check_struct struct{
    error_code string
    err error
}
func check_po_exist(t* purchase_order)(bool,error) {
    var get_po_no string
    var err error
    err=db.QueryRow("select purchase_order_id from t_purchase_order where associated_po_no=?",t.po_no).Scan(&get_po_no)
    if err!=nil{
        return false,err
    }else if get_po_no!=""{
        //修改purchase_order_id
        t.purchase_order_id=get_po_no
        return true,nil//存在po_no
    }
    return false,nil
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
    _,err:=os.Stat(configuration.Nfs_path+po_url)
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
            t=check_struct{error_deliver_notes_packing_method_id,errors.New(`deliver_notes packing_method_id missed`)}
        }
    }
    error_chan<- t
}
func check_deliver_notes_deliver_note_no(deliver_notes []Deliver_notes,error_chan chan<- check_struct) {
    // error_check_deliver_notes_deliver_note_no
    var t check_struct
    for _,d:=range deliver_notes{
        // fmt.Println(reflect.TypeOf(d))
        var note_id string
        db.QueryRow("select note_id from t_goods_delivery_note where associated_goods_delivery_note_no=?",d.Gdn_no).Scan(&note_id)
        if note_id!= ""{
            t=check_struct{error_check_deliver_notes_deliver_note_no,errors.New(`deliver_notes deliver_note_no already exists`)}
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
            t=check_struct{error_check_logistic_provider,errors.New(`Deliver_notes logistic_provider_basic_id missed`)}
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
func check_deliver_notes_commercial_invoice(path string) bool{
    _,err:=os.Stat(configuration.Nfs_path+path)
    if err!=nil{
        return false
    }
    return true
}
func check_deliver_notes_packing_list(path string) bool{
    _,err:=os.Stat(configuration.Nfs_path+path)
    if err!=nil{
        return false
    }
    return true
}
func check_deliver_notes_bill_of_lading(path string) bool{
    _,err:=os.Stat(configuration.Nfs_path+path)
    if err!=nil{
        return false
    }
    return true
}
func check_deliver_notes_associated_so(path string) bool{
    _,err:=os.Stat(configuration.Nfs_path+path)
    if err!=nil{
        return false
    }
    return true
}
////////////////following check for gdn
func check_po_exist_for_gdn(origi *DeliverGoodsForPO,error_chan chan<- check_struct){
    var get_po_no string
    var t check_struct
    var err error
    for _,d:= range origi.Data.Deliver_notes{
        err=db.QueryRow("select purchase_order_id from t_purchase_order where associated_po_no=?",d.Po_no).Scan(&get_po_no)
        if err!=nil{
            t=check_struct{error_check_po_exist_for_gdn,errors.New(`po_exist_for_gdn missed`)}
            error_chan<- t
            return
        }else if get_po_no==""{
            // return true,nil//存在po_no
            t=check_struct{error_check_po_exist_for_gdn,errors.New(`po_exist_for_gdn missed`)}
            error_chan<- t
            return
        }
    }
    error_chan<- t
}
func po_check_data(origi *PoData)(string,error) {
    // var all_error map[string]error
    error_chan:=make(chan check_struct)
    go check_request_system(origi.Request_system,error_chan)
    go check_bill_type(origi.Data.Purchase_order.Bill_type,error_chan)
    // go check_po_url(origi.Data.Purchase_order.Po_url,error_chan)
    go check_status(origi.Data.Purchase_order.Status,error_chan)
    go check_supplier(origi.Data.Purchase_order.Supplier,error_chan)
    // go check_packing_method(origi.Data.Deliver_notes,error_chan)
    // go check_logistic_provider(origi.Data.Deliver_notes,error_chan)
    go check_ship_via(origi.Data.Purchase_order.Ship_via,error_chan)
    go check_trade_term(origi.Data.Purchase_order.Trade_term,error_chan)
    // go check_deliver_notes_deliver_note_no(origi.Data.Deliver_notes,error_chan)
    for i:=0;i<6;i++{
        err:=<-error_chan
        // fmt.Println("104:",err.error_code,err.err)
        if err.err!=nil{
            return err.error_code,err.err
        }
    } 
    return "",nil
}
func gdn_check_data(origi *DeliverGoodsForPO)(string,error) {
    // var all_error map[string]error
    error_chan:=make(chan check_struct)
    go check_packing_method(origi.Data.Deliver_notes,error_chan)
    go check_logistic_provider(origi.Data.Deliver_notes,error_chan)
    go check_deliver_notes_deliver_note_no(origi.Data.Deliver_notes,error_chan)
    go check_po_exist_for_gdn(origi,error_chan)
    go check_request_system(origi.Request_system,error_chan)
    for i:=0;i<5;i++{
        err:=<-error_chan
        // fmt.Println("104:",err.error_code,err.err)
        if err.err!=nil{
            return err.error_code,err.err
        }
    } 
    return "",nil
}