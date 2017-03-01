 package main
 import (
    "time"
    "logger"
)
func insert_goods_delivery_note_attachment(po_no,file_name,url string)error {
    var err error
    _, err = db.Exec(
        `INSERT INTO t_goods_delivery_note_attachment(
        attachment_id,goods_delivery_note_id,file_name,language_id,sort_no,format,url,note,createAt,createBy,dr,data_version) 
        VALUES (?,?,?,?,?,?,?,?,?,?,?,?)`,
        rand_string(20),
        po_no,
        file_name,
        "language_id",
        1,
        0,
        url,
        "",
        time.Now().Format("2006-01-02 15:04:05"),
        "go_fcgi",
        0,
        1)
    return err
}
func insert_note_attachment(
    t *purchase_order,
    origi *DeliverGoodsForPO)error {
    var err error

    for _,d:= range origi.Data.Deliver_notes{
        err= insert_goods_delivery_note_attachment(t.po_no,d.Commercial_invoice.Ci_no,d.Commercial_invoice.Ci_url)
        
        err= insert_goods_delivery_note_attachment(t.po_no,d.Packing_list.Pl_no,d.Packing_list.Pl_url)
       
        err= insert_goods_delivery_note_attachment(t.po_no,d.Bill_of_lading.Bl_no,d.Bill_of_lading.Bl_url)
       
        err= insert_goods_delivery_note_attachment(t.po_no,d.Associated_so.Associated_so_no,d.Associated_so.Associated_so_url)

        if err!=nil{
            logger.Info("insert to goods_delivery_note_attachment:"+err.Error()) 
        }
    }
    
    
    return err
}