 package main
 import (
    
    "time"
)
func insert_purchase_order_detail(t *purchase_order,origi *DeliverGoodsForPO)error {
	var err error
	for _,detail:= range origi.Data.Purchase_order.Detail{
		_, err = db.Exec(
        `INSERT INTO t_purchase_order_detail(detail_id,purchase_order_id,
		item_master_id,unit_price,quantity,uom_id,sub_amount,warranty,
		comments,note,createAt,createBy,dr,data_version) 
		VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
		rand_string(20),
		t.purchase_order_id,
		"item_master_id",
		detail.Unit_price,
		detail.Quantity,
		"detail.Uom_id",
		detail.Sub_total,
		detail.Warranty,
		detail.Comments,
		detail.Note,
		time.Now().Format("2006-01-02 15:04:05"),
		"go_fcgi",
		0,
		1)
	}
	return err
}