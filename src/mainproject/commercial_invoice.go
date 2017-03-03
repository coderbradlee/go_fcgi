 package main
 import (
    "time"
    // "logger"
    "fmt"
)
 func get_company_id_chan(company_id_chan chan<- string,company string) {
    var company_id string
    db.QueryRow("select company_id from t_company where short_name=?",company).Scan(&company_id)
     company_id_chan<-company_id
 }
 // func get_purchase_order_id_chan(purchase_order_id_chan chan<- string,po_no string) {
 //    var purchase_order_id string
 //    db.QueryRow("select purchase_order_id from t_purchase_order where po_no=?",po_no).Scan(&purchase_order_id)
 //     purchase_order_id_chan<-purchase_order_id
 // }
func insert_ci(ci *Commercial_invoice,t *purchase_order,
    origi *DeliverGoodsForPO,sd *shared_data)error {
    var err error
    ////////////////////////////
    // company_id_chan :=make(chan string)
    // go get_company_id_chan(company_id_chan,origi.Data.Purchase_order.Company)
    // company_id:=<-company_id_chan
    //////////////////////
    // purchase_order_id_chan :=make(chan string)
    // go get_purchase_order_id_chan(purchase_order_id_chan,origi.Data.Purchase_order.Po_no)
    // purchase_order_id:=<-purchase_order_id_chan

/////////////////////////////////////////////////////////////////

    _, err = db.Exec(
        `INSERT INTO t_commercial_invoice(
        invoice_id,company_id,invoice_no,invoice_date,sales_order_id,
        purchase_order_id,outbound_note_id,status,process_type,
        payment_dead_line,payment_received,payment_due,shipping_cost_total,markup_total,tax_total,sub_total,grand_total,approvedBy,approvedAt,note,createAt,createBy,dr,data_version) 
        VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
        rand_string(20),
        t.company_id,
        ci.Ci_no,
        ci.Ci_date,
        "",//sales_order_id
        t.purchase_order_id,
        "",//outbound_note_id
        ci.Status,
        ci.Invoice_type,//pending
        "",//payment_dead_line//pending
        0,//payment_received//pending
        0,//payment_due//pending
        0,//shipping_cost_total,//pending
        0,//markup_total,//pending
        0,//tax_total,//pending
        0,//sub_total,//pending
        ci.Total_amount,
        ci.Approved_by,
        "",//pending approvedAt
        ci.Note,
        time.Now().Add(sd.company_time_zone).Format("2006-01-02 15:04:05"),
        "go_fcgi",
        0,
        1)
    fmt.Println("Commercial_invoice.go 61")
    return err
}

func insert_commercial_invoice(
    level3_chan chan<- error,
    t *purchase_order,
    origi *DeliverGoodsForPO,sd *shared_data) {
    var err error
   
    for _,d:= range origi.Data.Deliver_notes{
        err= insert_ci(&d.Commercial_invoice,t,origi,sd)
        if err!=nil{
            level3_chan<- err
        }   
    }
    level3_chan<-err
    fmt.Println("Commercial_invoice.go 78")
}