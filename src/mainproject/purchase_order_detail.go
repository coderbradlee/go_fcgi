 package main
 import (
    // "fmt"
    "time"
    "errors"
)

func insert_purchase_order_detail(t *purchase_order,origi *PoData,sd *shared_data)(string,error) {
	var err error
	system_account_id_chan :=make(chan string)
    go get_system_account_id_chan(system_account_id_chan,t.Data.Purchase_order.Created_by)
    created_by:=<-system_account_id_chan
	for _,detail:= range origi.Data.Purchase_order.Detail{
		// item_master_id:=get_item_master_id(detail.Item_no,detail.Product_name,detail.Product_code)
		// uom_id:=get_uom_id(detail.Uom)
		// fmt.Println(sd.company_time_zone)
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
            if uom_id==""{
            	return error_purchase_order_detail_uom_id,errors.New("purchase_order.detail uom_id is missed")
            }
            if item_master_id==""{
            	return error_purchase_order_detail_item_master_id,errors.New("purchase_order.detail item_master_id is missed")
            }
            
		_, err = db.Exec(
        `INSERT INTO t_purchase_order_detail(detail_id,purchase_order_id,
		item_master_id,unit_price,quantity,uom_id,amount,warranty,
		comments,note,createAt,createBy,updateBy,dr,data_version) 
		VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
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
		time.Now().Add(sd.company_time_zone).Format("2006-01-02 15:04:05"),
		created_by,
		"go_fcgi",
		0,
		1)
	}
	if err!=nil{
		return error_insert_purchase_order_detail,err
	}
	return "",nil
}