 package main
 import (
    "time"
    "logger"
)
 func get_bill_type_id_chan(bill_type_id_chan chan<- string,bill_type string) {
    var bill_type_id string
    db.QueryRow("select bill_type_id_chan from t_bill_type where name=?",bill_type).Scan(&bill_type_id)
     bill_type_id_chan<-bill_type_id
 }
func insert_goods_receipt(t *purchase_order,
    origi *DeliverGoodsForPO,sd *shared_data)error {
    var err error
    bill_type_id_chan :=make(chan string)
    go get_bill_type_id_chan(bill_type_id_chan,origi.Data.Purchase_order.Bill_type)
    bill_type_id:=<-bill_type_id_chan
    _, err = db.Exec(
        `INSERT INTO t_goods_delivery_note_detail(
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
            createAt,
            createBy,
            dr,
            data_version) 
        VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
        rand_string(20),
        sd.goods_receipt_no,
        bill_type_id,
        company_id,
        purchase_order_id,
        sd.goods_delivery_note_id,
        status,
        receipt_date,
        from_system_code,
        approved_by,
        approved_at,
        time.Now().Add(sd.company_time_zone).Format("2006-01-02 15:04:05"),
        "go_fcgi",
        0,
        1)
    return err
}

