 package main
 import (
    "time"
    // "logger"
    // "fmt"
    "errors"
)

func insert_ci(d *Deliver_notes,sd *shared_data)(string,error) {
    var err error
    ci:=d.Commercial_invoice
    if ci.Status!=1{
        return error_commercial_invoice_status,errors.New("commercial_invoice.status!=1")
    }
    
    company_id_chan :=make(chan string)
    go get_company_id_chan(company_id_chan,ci.Company)
    company_id:=<-company_id_chan

    _, err = db.Exec(
        `INSERT INTO t_commercial_invoice(
        invoice_id,company_id,associated_invoice_no,associated_system_code,invoice_no,invoice_date,type,sales_order_id,outbound_note_id,status,process_type,
        payment_dead_line,payment_due,shipping_cost_total,markup_total,tax_total,sub_total,grand_total,url,approvedBy,approvedAt,note,createAt,createBy,dr,data_version,varchar_1) 
        VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
        rand_string(20),
        company_id,
        ci.Ci_no,
        1,
        "invoice_no",//need to get from redis flowno
        ci.Ci_date,
        1,
        "",//sales_order_id
        "",//outbound_note_id
        ci.Status,
        ci.Invoice_type,//pending
        "",//payment_dead_line//pending
        0,//payment_due//pending
        0,//shipping_cost_total,//pending
        0,//markup_total,//pending
        0,//tax_total,//pending
        0,//sub_total,//pending
        ci.Total_amount,
        ci.Ci_url,
        ci.Approved_by,
        ci.Ci_date,//pending approvedAt
        ci.Note,
        time.Now().Add(sd.company_time_zone).Format("2006-01-02 15:04:05"),
        ci.Created_by+" go_fcgi",
        0,
        1,sd.goods_delivery_note_id)
    // fmt.Println("ci")
    return error_insert_commercial_invoice,err
}

func insert_commercial_invoice(
    d *Deliver_notes,sd *shared_data,note_id string)(string,error) {
    var err error
    var s string
    // for _,d:= range origi.Data.Deliver_notes{
        s,err= insert_ci(d,sd)
        if err!=nil{
            return s,err
        }   
    // }
    return error_insert_commercial_invoice,err
}