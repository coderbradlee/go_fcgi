 package main
 import (
    "time"
    "logger"
)
func insert_goods_delivery_note_attachment(po_no,file_name,url,language string,sort_no int)error {
    var err error
    _, err = db.Exec(
        `INSERT INTO t_goods_delivery_note_attachment(
        attachment_id,goods_delivery_note_id,file_name,language_id,sort_no,format,url,note,createAt,createBy,dr,data_version) 
        VALUES (?,?,?,?,?,?,?,?,?,?,?,?)`,
        rand_string(20),
        po_no,
        file_name,
        language,
        sort_no,
        0,
        url,
        "",
        time.Now().Format("2006-01-02 15:04:05"),
        "go_fcgi",
        0,
        1)
    return err
}
func get_language_id(company string)string{
    var language_id string
    db.QueryRow("select default_language_id from t_company where short_name=?",company).Scan(&language_id)
    return language_id
}
func get_sort_no()int {
    var sortno int 
    db.QueryRow("select sort_no from t_goods_delivery_note_attachment ORDER BY sort_no desc LIMIT 1").Scan(&sortno)
    return sortno
}
func insert_note_attachment(
    t *purchase_order,
    origi *DeliverGoodsForPO)error {
    var err error
    language:=get_language_id(origi.Data.Purchase_order.Company)
    sort_no:=get_sort_no()
    for _,d:= range origi.Data.Deliver_notes{
        err= insert_goods_delivery_note_attachment(t.po_no,d.Commercial_invoice.Ci_no,d.Commercial_invoice.Ci_url,language,sort_no+1)
        
        err= insert_goods_delivery_note_attachment(t.po_no,d.Packing_list.Pl_no,d.Packing_list.Pl_url,language,sort_no+2)
       
        err= insert_goods_delivery_note_attachment(t.po_no,d.Bill_of_lading.Bl_no,d.Bill_of_lading.Bl_url,language,sort_no+3)
       
        err= insert_goods_delivery_note_attachment(t.po_no,d.Associated_so.Associated_so_no,d.Associated_so.Associated_so_url,language,sort_no+4)

        if err!=nil{
            logger.Info("insert to goods_delivery_note_attachment:"+err.Error()) 
        }
    }
    
    
    return err
}