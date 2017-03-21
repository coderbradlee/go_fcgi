 package main
 import (
    "time"
    "logger"
)
func insert_goods_delivery_note_detail(detail *Deliver_notes_detail,item_master_id,uom_id string,sd *shared_data)(string,error) {
    var err error
    _, err = db.Exec(
        `INSERT INTO t_goods_delivery_note_detail(
        detail_id,
        goods_delivery_note_id,
        item_master_id,
        delivery_qty,
        uom_id,
        unit_price,
        currency_id,
        amount,
        note,
        createAt,
        createBy,
        dr,
        data_version) 
        VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?)`,
        rand_string(20),
        sd.goods_delivery_note_id,
        item_master_id,
        detail.Quantity,
        uom_id,
        detail.Unit_price,
        "currency_id",
        detail.Sub_total,
        "note"
        time.Now().Add(sd.company_time_zone).Format("2006-01-02 15:04:05"),
        "go_fcgi",
        0,
        1)
    return error_insert_goods_delivery_note_detail,err
}

func insert_note_detail(
    origi *DeliverGoodsForPO,sd *shared_data)(string,error) {
    var err error
    var s string
    for _,d:= range origi.Data.Deliver_notes{
        for _,detail:=range d.Detail{
            // item_master_id:=get_item_master_id(detail.Item_no,detail.Product_name,detail.Product_code)
            item_master_id_chan :=make(chan string)
            go get_item_master_id_chan(item_master_id_chan,detail.Item_no,detail.Product_name,detail.Product_code)
            // item_master_id:=<-item_master_id_chan
            ////////////////////////////////////////
            // uom_id:=get_uom_id(detail.Uom)

            uom_id_chan :=make(chan string)
            go get_uom_id_chan(uom_id_chan,detail.Uom)
            uom_id:=<-uom_id_chan
            item_master_id:=<-item_master_id_chan
            
            s,err= insert_goods_delivery_note_detail(&detail,item_master_id,uom_id,sd)
            if err!=nil{
                logger.Info("insert to insert_goods_delivery_note_detail:"+err.Error()) 
                return s,err
            }
        }
    }    
    return error_insert_goods_delivery_note_detail,err
}