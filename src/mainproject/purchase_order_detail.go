 package main
 import (
    
    "time"
)
func get_item_master_id(item_no,product_name,product_code string) string{
	var item_basic_id string
    err := db.QueryRow("select item_basic_id from t_item_basic where item_no=?",item_no).Scan(&item_basic_id)

	var item_master_id string
    err := db.QueryRow("select item_master_id from t_item_master where item_basic_id=? and product_code=? and product_name=?",item_basic_id,product_code,product_name).Scan(&item_master_id)

    return item_master_id
}
func get_uom_id(uom string) string{
	var uom_id string
    err := db.QueryRow("select uom_id from t_uom where uom=?",uom).Scan(&uom_id)

    return uom_id
}
func insert_purchase_order_detail(t *purchase_order,origi *DeliverGoodsForPO)error {
	var err error
	for _,detail:= range origi.Data.Purchase_order.Detail{
		item_master_id:=get_item_master_id(detail.Item_no,detail.Product_name,detail.Product_code)
		uom_id:=get_uom_id(detail.Uom)
		_, err = db.Exec(
        `INSERT INTO t_purchase_order_detail(detail_id,purchase_order_id,
		item_master_id,unit_price,quantity,uom_id,sub_amount,warranty,
		comments,note,createAt,createBy,dr,data_version) 
		VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
		rand_string(20),
		t.purchase_order_id,
		item_master_id,
		detail.Unit_price,
		detail.Quantity,
		uom_id,
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