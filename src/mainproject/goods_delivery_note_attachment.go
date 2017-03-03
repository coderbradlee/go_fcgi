 package main
 import (
    "time"
    "logger"
)
func insert_goods_delivery_note_attachment(file_name,url,language string,sort_no int,sd *shared_data)error {
    var err error
    _, err = db.Exec(
        `INSERT INTO t_goods_delivery_note_attachment(
        attachment_id,goods_delivery_note_id,file_name,language_id,sort_no,format,url,note,createAt,createBy,dr,data_version) 
        VALUES (?,?,?,?,?,?,?,?,?,?,?,?)`,
        rand_string(20),
        sd.goods_delivery_note_id,
        file_name,
        language,
        sort_no,
        0,
        url,
        "",
        time.Now().Add(sd.company_time_zone).Format("2006-01-02 15:04:05"),
        "go_fcgi",
        0,
        1)
    return err
}
func get_language_id_chan(language_id_chan chan<- string,company string){
    var language_id string
    db.QueryRow("select default_language_id from t_company where short_name=?",company).Scan(&language_id)
    language_id_chan<- language_id
}
func get_sort_no_chan(sort_no_chan chan<- int) {
    var sortno int 
    db.QueryRow("select sort_no from t_goods_delivery_note_attachment ORDER BY sort_no desc LIMIT 1").Scan(&sortno)
    sort_no_chan<-sortno
}
func insert_note_attachment(
    level4_chan chan<- error,
    t *purchase_order,
    origi *DeliverGoodsForPO,sd *shared_data) {
    var err error
    // language:=get_language_id(origi.Data.Purchase_order.Company)
    language_chan :=make(chan string)
    go get_language_id_chan(language_chan,origi.Data.Purchase_order.Company)
    // language:=<-language_chan
/////////////////////////////////////////////
    // sort_no:=get_sort_no()
    sort_no_chan :=make(chan int)
    go get_sort_no_chan(sort_no_chan)
    sort_no:=<-sort_no_chan
    language:=<-language_chan

    for _,d:= range origi.Data.Deliver_notes{
        err= insert_goods_delivery_note_attachment(d.Commercial_invoice.Ci_no,d.Commercial_invoice.Ci_url,language,sort_no+1,sd)
        
        err= insert_goods_delivery_note_attachment(d.Packing_list.Pl_no,d.Packing_list.Pl_url,language,sort_no+2,sd)
       
        err= insert_goods_delivery_note_attachment(d.Bill_of_lading.Bl_no,d.Bill_of_lading.Bl_url,language,sort_no+3,sd)
       
        err= insert_goods_delivery_note_attachment(d.Associated_so.Associated_so_no,d.Associated_so.Associated_so_url,language,sort_no+4,sd)

        if err!=nil{
            logger.Info("insert to goods_delivery_note_attachment:"+err.Error()) 
        }
    }
    level4_chan<- err
}