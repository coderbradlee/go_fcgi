 package main
 import (
    
    "time"
)
func insert_purchase_order_detail(t *purchase_order,origi *DeliverGoodsForPO)error {
	for index,detail:= range origi.Data.Purchase_order.detail{
		_, err := db.Exec(
        `INSERT INTO t_purchase_order_detail(detail_id,purchase_order_id,
		item_master_id,unit_price,quantity,uom_id,sub_amount,warranty,
		comments,note,createAt,createBy,dr,data_version) 
		VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?.?)`,
		rand_string(20),
		t.purchase_order_id,
		"item_master_id",
		detail.unit_price,
		detail.quantity,
		"detail.uom_id",
		detail.sub_amount,
		detail.warranty,
		detail.comments,
		detail.note,
		time.Now().Format("2006-01-02 15:04:05"),
		"go_fcgi",
		0,
		1)
	}
	return err
}