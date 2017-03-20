 package main
 import (
    "time"
    // "logger"
    "fmt"
    "errors"
)

func insert_ci(ci *Commercial_invoice,
    origi *DeliverGoodsForPO,sd *shared_data)(string,error) {
    var err error
    
    if ci.Status!=1{
        return error_commercial_invoice_status,errors.New("commercial_invoice.status!=1")
    }
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
        ci.Created_by+" go_fcgi",
        0,
        1)
    // fmt.Println("ci")
    return error_insert_commercial_invoice,err
}

func insert_commercial_invoice(
    origi *DeliverGoodsForPO,sd *shared_data)(string,error) {
    var err error
    var s string
    for _,d:= range origi.Data.Deliver_notes{
        s,err= insert_ci(&d.Commercial_invoice,origi,sd)
        if err!=nil{
            return s,err
        }   
    }
    return error_insert_commercial_invoice,err
}