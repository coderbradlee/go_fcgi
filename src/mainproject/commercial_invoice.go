 package main
 import (
    // "time"
    // "logger"
    "fmt"
)
func insert_ci(ci *Commercial_invoice)error {
    var err error
    // _, err = db.Exec(
    //     `INSERT INTO t_goods_delivery_note_attachment(
    //     attachment_id,goods_delivery_note_id,file_name,language_id,sort_no,format,url,note,createAt,createBy,dr,data_version) 
    //     VALUES (?,?,?,?,?,?,?,?,?,?,?,?)`,
    //     rand_string(20),
    //     po_no,
    //     file_name,
    //     language,
    //     sort_no,
    //     0,
    //     url,
    //     "",
    //     time.Now().Format("2006-01-02 15:04:05"),
    //     "go_fcgi",
    //     0,
    //     1)
    fmt.Println("ci")
    return err
}

func insert_commercial_invoice(
    t *purchase_order,
    origi *DeliverGoodsForPO)error {
    var err error
   
    for _,d:= range origi.Data.Deliver_notes{
        err= insert_ci(&d.Commercial_invoice)
        if err!=nil{
            return err
        }   
    }
    return err
}