 package main
 import (
    "time"
    "logger"
    "fmt"
    "net/http"
    "io/ioutil"
    "strconv"
    "encoding/json"
    "errors"
)

// func get_bill_type_id()string {
//     var bill_type_id string
//     db.QueryRow("select bill_type_id from t_bill_type where code='GDN'").Scan(&bill_type_id)
//     return bill_type_id
// }
// func get_bill_type_id_chan(company_id_chan chan<- string,company string) {
//     var company_id string
//     db.QueryRow("select company_id from t_company where short_name=?",company).Scan(&company_id)
//      company_id_chan<-company_id
//  }
// func get_vendor_master_id(vendor_basic_id string)string {
//     var vendor_master_id string
//     db.QueryRow("select vendor_master_id from t_vendor_master where vendor_basic_id=?",vendor_basic_id).Scan(&vendor_master_id)
//     return vendor_master_id
// }
func get_vendor_master_id_chan(vendor_master_id_chan chan<- string,vendor_basic_id string) {
    var vendor_master_id string
    db.QueryRow("select vendor_master_id from t_vendor_master where vendor_basic_id=?",vendor_basic_id).Scan(&vendor_master_id)
    vendor_master_id_chan<-vendor_master_id
}
// func get_trade_term_id(Trade_term string)string {
//     var trade_term_id string
//     db.QueryRow("select trade_term_id from t_trade_term where short_name=?",Trade_term).Scan(&trade_term_id)
//     return trade_term_id
// }

func get_trade_term_id_chan(trade_term_id_chan chan<- string,Trade_term string) {
    var trade_term_id string
    db.QueryRow("select trade_term_id from t_trade_term where short_name=?",Trade_term).Scan(&trade_term_id)
    trade_term_id_chan<-trade_term_id
}
func get_buyer_id_chan(buyer_id_chan chan<- string,buyer string) {
    var buyer_id string
    db.QueryRow("select company_id from t_company where short_name=?",buyer).Scan(&buyer_id)
    buyer_id_chan<-buyer_id
}

func get_transport_term_id_chan(transport_term_id_chan chan<- string,ship_via string) {
    var transport_term_id string
    db.QueryRow("select ship_via_id from t_ship_via where full_name=?",ship_via).Scan(&transport_term_id)
    transport_term_id_chan<-transport_term_id
}
func get_packing_method_id_chan(packing_method_id_chan chan<- string,Packing_method string) {
    var packing_method_id string
    db.QueryRow("select packing_method_id from t_packing_method where name=?",Packing_method).Scan(&packing_method_id)
    packing_method_id_chan<- packing_method_id
}
func get_logistic_master_id_chan(logistic_master_id_chan chan<- string,Logistic string) {
    var logistic_master_id string
    db.QueryRow("select logistic_provider_master_id from t_logistic_provider_master where native_name=?",Logistic).Scan(&logistic_master_id)
    logistic_master_id_chan<-logistic_master_id
}
func get_logistic_contact_id_chan(logistic_contact_id_chan chan<- string,Logistic_contact string) {
    var logistic_contact_id string
    db.QueryRow("select logistic_contact_id from t_logistic_provider_master where native_name=?",Logistic_contact).Scan(&logistic_contact_id)
    logistic_contact_id_chan<-logistic_contact_id
}
func get_currency_id(currency_id_chan chan<- string,currency string) {
    var currency_id string
    db.QueryRow("select currency_id from t_currency where code=?",currency).Scan(&currency_id)
    currency_id_chan<-currency_id
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
    i, err3 := strconv.Atoi(data.FlowNo)
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
    // flow,err:=get_flow_no(short)
    // if err!=nil{
    //     return "",err
    // }
    flow:="000001"
    goods_delivery_note_no+=flow//get int,format to 6bit,then convert to string
    return goods_delivery_note_no,nil
}
//为了在response中回传发货号，设置全局变量goods_receipt_no
// var goods_receipt_no string
// var goods_delivery_note_id string
func insert_goods_delivery_note(t *purchase_order,origi *DeliverGoodsForPO,sd *shared_data)(string,error) {
    var err error
    // log.Println("insert_goods_delivery_note")
    for _,deliver_notes:= range origi.Data.Deliver_notes{
        // bill_type_id:=get_bill_type_id(t.Bill_type)
        // bill_type_id:=get_bill_type_id()

        bill_type_id_chan :=make(chan string)
        go get_bill_type_id_chan(bill_type_id_chan,origi.Data.Purchase_order.Bill_type)
        // bill_type_id:=<-bill_type_id_chan
        /////////////////////////////////////////////
        // vendor_master_id:=get_vendor_master_id(t.vendor_basic_id)
        vendor_master_id_chan :=make(chan string)
        go get_vendor_master_id_chan(vendor_master_id_chan,t.vendor_basic_id)
        // vendor_master_id:=<-vendor_master_id_chan
////////////////////////////////////////////////////////////
// trade_term_id:=get_trade_term_id(deliver_notes.Trade_term)
        trade_term_id_chan :=make(chan string)
        go get_trade_term_id_chan(trade_term_id_chan,deliver_notes.Trade_term)
        // trade_term_id:=<-trade_term_id_chan

///////////////////////////////////////////////////////////
        // buyer_id:=get_buyer_id(deliver_notes.Buyer)
        buyer_id_chan :=make(chan string)
        go get_buyer_id_chan(buyer_id_chan,deliver_notes.Buyer)
        // buyer_id:=<-buyer_id_chan
///////////////////////////////////////////////////
// transport_term_id:=get_transport_term_id(deliver_notes.Ship_via)
        transport_term_id_chan :=make(chan string)
        go get_transport_term_id_chan(transport_term_id_chan,deliver_notes.Ship_via)
        // transport_term_id:=<-transport_term_id_chan
/////////////////////////////////////////////////////////////////         
        //packing_method_id:=get_packing_method_id(deliver_notes.Packing_method)
        packing_method_id_chan :=make(chan string)
        go get_packing_method_id_chan(packing_method_id_chan,deliver_notes.Packing_method)
        // packing_method_id:=<-packing_method_id_chan
//////////////////////////////////////////////////////////////
        // logistic_master_id:=get_logistic_master_id(deliver_notes.Logistic)
        logistic_master_id_chan :=make(chan string)
        go get_logistic_master_id_chan(logistic_master_id_chan,deliver_notes.Logistic)
        // logistic_master_id:=<-logistic_master_id_chan
////////////////////////////////////////////////////////////////////
        // logistic_contact_id:=get_logistic_contact_id(deliver_notes.Logistic_contact)
        logistic_contact_id_chan :=make(chan string)
        go get_logistic_contact_id_chan(logistic_contact_id_chan,deliver_notes.Logistic_contact)
        // logistic_contact_id:=<-logistic_contact_id_chan
        ////////////////////////////////////////////////////
        // currency_id_chan :=make(chan string)
        // go get_currency_id(currency_id_chan,deliver_notes.Currency)
    // t_purchase_order.currency_id=<-currency_id_chan
/////////////////////////////////////////////////////////////////////// 
///     在这里集中同步
        logistic_master_id:=<-logistic_master_id_chan
        packing_method_id:=<-packing_method_id_chan
        transport_term_id:=<-transport_term_id_chan
        buyer_id:=<-buyer_id_chan
        trade_term_id:=<-trade_term_id_chan
        vendor_master_id:=<-vendor_master_id_chan
        bill_type_id:=<-bill_type_id_chan
        logistic_contact_id:=<-logistic_contact_id_chan
        var exist bool
        exist=check_deliver_notes_commercial_invoice(deliver_notes.Commercial_invoice.Ci_url)
        if !exist{
             return error_check_deliver_notes_commercial_invoice,errors.New("deliver_notes commercial_invoice file is missed")
        }
        exist=check_deliver_notes_packing_list(deliver_notes.Packing_list.Pl_url)
        if !exist{
             return error_check_deliver_notes_packing_list,errors.New("deliver_notes packing_list file is missed")
        }
        exist=check_deliver_notes_bill_of_lading(deliver_notes.Bill_of_lading.Bl_url)
        if !exist{
             return error_check_deliver_notes_bill_of_lading,errors.New("deliver_notes bill_of_lading file is missed")
        }
        exist=check_deliver_notes_associated_so(deliver_notes.Associated_so.Associated_so_url)
        if !exist{
             return error_check_deliver_notes_associated_so,errors.New("deliver_notes associated_so file is missed")
        }


        if transport_term_id==""{
            return error_deliver_notes_transport_term_id,errors.New("deliver_notes transport_term_id is missed")
        }
        if buyer_id==""{
            //return error_buyer_id,errors.New("buyer_id is missed")
        }
        if trade_term_id==""{
            return error_deliver_notes_trade_term_id,errors.New("deliver_notes trade_term_id is missed")
        }
        if vendor_master_id==""{
            return error_deliver_notes_vendor_master_id,errors.New("deliver_notes vendor_master_id is missed")
        }
        if bill_type_id==""{
            // return error_bill_type_id,errors.New("bill_type_id is missed")
        }
        ///////////////////////////////////////////////
        goods_delivery_note_no,err:=get_goods_delivery_note_no(origi.Data.Purchase_order.Company)
        sd.goods_receipt_no=goods_delivery_note_no
        if err!=nil{
            return "",err
        }
        sd.goods_delivery_note_id=rand_string(20)
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
        sd.goods_delivery_note_id,
        goods_delivery_note_no,//goods_delivery_note_no 待定
        bill_type_id,
        t.company_id,
        t.purchase_order_id,
        buyer_id,
        vendor_master_id,
        0,
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
        t.contact_account_id,//receiver 待定
        deliver_notes.Total_freight_charges,
        deliver_notes.Total_insurance_fee,
        deliver_notes.Total_excluded_tax,
        "",//note
        time.Now().Add(sd.company_time_zone).Format("2006-01-02 15:04:05"),
        "go_fcgi",
        0,
        1)
    }
    if err!=nil{
        logger.Info("insert_goods_delivery_note:"+err.Error()) 
        return error_insert_goods_delivery_note,err
    }
    return "",err
}