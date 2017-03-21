 package main
 import (
    "time"
    "logger"
    "fmt"
    // "net/http"
    // "io/ioutil"
    // "strconv"
    // "encoding/json"
    "errors"
    // "errgroup"
)

//为了在response中回传发货号，设置全局变量goods_receipt_no
// var goods_receipt_no string
// var goods_delivery_note_id string
func insert_goods_delivery_note(origi *DeliverGoodsForPO,sd *shared_data)(string,error) {
    var err error
    
    for _,deliver_notes:= range origi.Data.Deliver_notes{
        // bill_type_id:=get_bill_type_id(t.Bill_type)
        // bill_type_id:=get_bill_type_id()
        fmt.Println("insert_goods_delivery_note")

        fmt.Println("Bill_type:",deliver_notes.Bill_type)
        bill_type_id_chan :=make(chan string)
        go get_bill_type_id_chan(bill_type_id_chan,deliver_notes.Bill_type)
        //bill_type_id:=<-bill_type_id_chan
        
////////////////////////////////////////////////////////////
// trade_term_id:=get_trade_term_id(deliver_notes.Trade_term)
        trade_term_id_chan :=make(chan string)
        go get_trade_term_id_chan(trade_term_id_chan,deliver_notes.Trade_term)
        //trade_term_id:=<-trade_term_id_chan

///////////////////////////////////////////////////////////
        // buyer_id:=get_buyer_id(deliver_notes.Buyer)
        buyer_id_chan :=make(chan string)
        go get_buyer_id_chan(buyer_id_chan,deliver_notes.Buyer)
        //buyer_id:=<-buyer_id_chan
///////////////////////////////////////////////////
// transport_term_id:=get_transport_term_id(deliver_notes.Ship_via)
        transport_term_id_chan :=make(chan string)
        go get_transport_term_id_chan(transport_term_id_chan,deliver_notes.Ship_via)
        //transport_term_id:=<-transport_term_id_chan
/////////////////////////////////////////////////////////////////         
        //packing_method_id:=get_packing_method_id(deliver_notes.Packing_method)
        packing_method_id_chan :=make(chan string)
        go get_packing_method_id_chan(packing_method_id_chan,deliver_notes.Packing_method)
        //packing_method_id:=<-packing_method_id_chan
//////////////////////////////////////////////////////////////
        // logistic_master_id:=get_logistic_master_id(deliver_notes.Logistic)
        logistic_master_id_chan :=make(chan string)
        go get_logistic_master_id_chan(logistic_master_id_chan,deliver_notes.Logistic)
        //logistic_master_id:=<-logistic_master_id_chan
////////////////////////////////////////////////////////////////////
        // logistic_contact_id:=get_logistic_contact_id(deliver_notes.Logistic_contact)
        logistic_contact_id_chan :=make(chan string)
        go get_logistic_contact_id_chan(logistic_contact_id_chan,deliver_notes.Logistic_contact)
        //logistic_contact_id:=<-logistic_contact_id_chan
        ////////////////////////////////////////////////////
        // currency_id_chan :=make(chan string)
        // go get_currency_id(currency_id_chan,deliver_notes.Currency)
    // t_purchase_order.currency_id=<-currency_id_chan
        purchase_order_table_chan :=make(chan purchase_order_part)
        go get_purchase_order_table_chan(purchase_order_table_chan,deliver_notes.Po_no)
        purchase_order_table:=<-purchase_order_table_chan//get_vendor_master_id用到所以需要提前拿到
        /////////////////////////////////////////////
        vendor_master_id_chan :=make(chan string)
        go get_vendor_master_id_chan(vendor_master_id_chan,purchase_order_table.vendor_basic_id)
        // vendor_master_id:=<-vendor_master_id_chan                 
        // ///////////////////
        export_country_id_chan :=make(chan string)
        go get_country_id_chan(export_country_id_chan,deliver_notes.Export_country)
        import_country_id_chan :=make(chan string)
        go get_country_id_chan(import_country_id_chan,deliver_notes.Import_country)    
        //////////////////////////////////////////////   
        loading_port_id_chan :=make(chan string)
        go get_port_id_chan(loading_port_id_chan,deliver_notes.Loading_port)
        unloading_port_id_chan :=make(chan string)
        go get_port_id_chan(unloading_port_id_chan,deliver_notes.Unloading_port)                        
/////////////////////////////////////////////////////////////////////// 
///     在这里集中同步
        logistic_master_id:=<-logistic_master_id_chan
        packing_method_id:=<-packing_method_id_chan
        transport_term_id:=<-transport_term_id_chan
        buyer_id:=<-buyer_id_chan
        trade_term_id:=<-trade_term_id_chan
        vendor_master_id:=<-vendor_master_id_chan
        // vendor_master_id:="vendor_master_id"
        bill_type_id:=<-bill_type_id_chan
        logistic_contact_id:=<-logistic_contact_id_chan
        export_country_id:=<-export_country_id_chan   
        import_country_id:=<-import_country_id_chan 
        loading_port_id:=<-loading_port_id_chan
        unloading_port_id:=<-unloading_port_id_chan

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
        // goods_delivery_note_no,err:=get_goods_delivery_note_no(origi.Data.Purchase_order.Company)
        // goods_delivery_note_no,err:=get_goods_delivery_note_no(deliver_notes.Deliver_note_no)
        // sd.goods_receipt_no=goods_delivery_note_no
        // if err!=nil{
        //     return "",err
        // }
        goods_delivery_note_no:=deliver_notes.Gdn_no
        sd.goods_delivery_note_id=rand_string(20)
        _, err = db.Exec(
        `INSERT INTO t_goods_delivery_note(
        note_id,goods_delivery_note_no,bill_type_id,company_id,
        purchase_order_id,buyer_id,vendor_master_id,fulfill_status,
        export_country_id,loading_port_id,import_country_id,unloading_port_id,trade_term_id,ship_via_id,packing_method_id,
        logistic_provider_master_id,logistic_provider_contact_id,etd,
        eta,atd,ata,customs_clearance_date,receiver,total_freight_charges,
        total_insurance_fee,total_excluded_tax,picking_list_no,picking_list_url,bill_of_lading_no,bill_of_lading_url,associated_so_no,associated_so_url,note,createAt,createBy,dr,
        data_version) 
        VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
        sd.goods_delivery_note_id,
        goods_delivery_note_no,//goods_delivery_note_no 待定
        bill_type_id,
        purchase_order_table.company_id,
        purchase_order_table.purchase_order_id,
        buyer_id,
        vendor_master_id,
        0,
        export_country_id,
        loading_port_id,
        import_country_id,
        unloading_port_id,
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
        deliver_notes.Supplier,//receiver 待定
        deliver_notes.Total_freight_charges,
        deliver_notes.Total_insurance_fee,
        deliver_notes.Total_excluded_tax,
        deliver_notes.Packing_list.Pl_no,
        deliver_notes.Packing_list.Pl_url,
        deliver_notes.Bill_of_lading.Bl_no,
        deliver_notes.Bill_of_lading.Bl_url,
        deliver_notes.Associated_so.Associated_so_no,
        deliver_notes.Associated_so.Associated_so_url,

        "",//note
        time.Now().Add(sd.company_time_zone).Format("2006-01-02 15:04:05"),
        "go_fcgi",
        0,
        1)
    if err!=nil{
        logger.Info("insert_goods_delivery_note:"+err.Error()) 
        return error_insert_goods_delivery_note,err
    }else{
        s,err:= insert_note_detail(&deliver_notes,sd)
        if err!=nil{
            logger.Info(s+":"+err.Error()) 
            return s,err
        }
    }
}
    return "",nil
}