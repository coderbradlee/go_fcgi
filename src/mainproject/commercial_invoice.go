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
    ///////////////////////////////////
    company_short_name_chan :=make(chan string)
    go get_company_short_name_chan(company_short_name_chan,ci.Company)
    company_short_name:=<-company_short_name_chan

    if company_short_name==""||len(company_short_name)>3{
        return error_insert_commercial_invoice,errors.New("get_company_short_name error")
    }
/////////////////////////////////////////
    system_account_id_chan :=make(chan string)
    go get_system_account_id_chan(system_account_id_chan,ci.Created_by)
    createBy:=<-system_account_id_chan
    ////////////////////////////////////
    approvedBy_chan :=make(chan string)
    go get_system_account_id_chan(approvedBy_chan,ci.Approved_by)
    approved_by:=<-approvedBy_chan
    //////////////////////////////////////////
    flow_no,err:=get_flow_no(company_short_name,"CI")
    if flow_no==""{
        return error_insert_commercial_invoice,errors.New("get_flow_no error")
    }
    invoice_no:="CI-"+company_short_name+"-"+time.Now().Format("20060102")+"-"+flow_no
    _, err = db.Exec(
        `INSERT INTO t_commercial_invoice(
        invoice_id,company_id,associated_invoice_no,associated_system_code,invoice_no,invoice_date,type,sales_order_id,outbound_note_id,status,process_type,
        payment_dead_line,payment_due,shipping_cost_total,markup_total,tax_total,sub_total,grand_total,url,approvedBy,approvedAt,note,createAt,createBy,updateBy,dr,data_version,varchar_1) 
        VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
        rand_string(20),
        company_id,
        ci.Ci_no,
        1,
        invoice_no,//need to get from redis flowno
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
        approved_by,
        ci.Ci_date,//pending approvedAt
        ci.Note,
        time.Now().Add(sd.company_time_zone).Format("2006-01-02 15:04:05"),
        createBy,
        "test_go_fcgi",
        0,
        1,sd.goods_delivery_note_id)
    // fmt.Println("ci")
    return error_insert_commercial_invoice,err
}

func insert_commercial_invoice(
    d *Deliver_notes,sd *shared_data)(string,error) {
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