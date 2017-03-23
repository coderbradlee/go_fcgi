 package main
 import (
    "time"
    // "logger"
)

func insert_goods_receipt(
    d *Deliver_notes,sd *shared_data)(string,error) {
    var err error
    // for _,d:= range origi.Data.Deliver_notes{
        bill_type_id_chan :=make(chan string)
        go get_bill_type_id_chan(bill_type_id_chan,d.Bill_type)
        bill_type_id:=<-bill_type_id_chan
        ///////////////////////////////
        company_id_chan :=make(chan string)
        go get_company_id_chan(company_id_chan,d.Company)
        company_id:=<-company_id_chan
////////////////////////////////////////////////
        purchase_order_id_chan :=make(chan string)
        go get_purchase_order_id_chan(purchase_order_id_chan,d.Po_no)
        purchase_order_id:=<-purchase_order_id_chan

    _, err = db.Exec(
        `INSERT INTO t_goods_receipt(
            receipt_id,
            goods_receipt_no,
            bill_type_id,
            company_id,
            purchase_order_id,
            goods_delivery_note_id,
            status,
            receipt_date,
            from_system_code,
            approved_by,
            approved_at,
            note,
            createAt,
            createBy,
            updateBy,
            dr,
            data_version) 
        VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
        rand_string(20),
        sd.goods_receipt_no,
        bill_type_id,
        company_id,
        purchase_order_id,
        sd.goods_delivery_note_id,
        0,//status
        "",//receipt_date,
        2,//from_system_code,
        d.Approved_by,//approved_by,
        "",//approved_at,
        d.Note,//note
        time.Now().Add(sd.company_time_zone).Format("2006-01-02 15:04:05"),
        d.Created_by,
        "go_fcgi",
        0,
        1)
    // }
    return error_insert_goods_receipt,err
}

