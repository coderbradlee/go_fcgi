 package main
 import (
	"errors"
	// "strings"
)
func get_company_id(company string) string{
	// var item_basic_id string
 //    db.QueryRow("select item_basic_id from t_item_basic where item_no=?",item_no).Scan(&item_basic_id)

	var company_id string
    db.QueryRow("select company_id from t_company where short_name=?",company).Scan(&company_id)

    return company_id
}
func get_shipping_method_id(Ship_via string) string{
	//cannot find the way to shipping_method_id
	// var shipping_method_id string
 //    db.QueryRow("select shipping_method_id from t_company where short_name=?",company).Scan(&shipping_method_id)
	var shipping_via_id string
    db.QueryRow("select ship_via_id from t_ship_via where full_name=?",Ship_via).Scan(&shipping_via_id)
    return shipping_via_id
}
func get_vendor_basic_id(supplier string)string {
	var vendor_basic_id string
    db.QueryRow("select vendor_basic_id from t_vendor_basic where short_name=?",supplier).Scan(&vendor_basic_id)
    return vendor_basic_id
}
func get_contact_account_id(vendor_basic_id string)string {
	var vendor_contact_id string
    db.QueryRow("select vendor_contact_id from t_vendor_contact where vendor_basic_id=?",vendor_basic_id).Scan(&vendor_contact_id)
    return vendor_contact_id
}
func check_po_exist(po_no string)error {
	var get_po_no string
    db.QueryRow("select po_no from t_purchase_order where po_no=?",po_no).Scan(&get_po_no)
    if get_po_no!=""{
    	return errors.New("po_no")
    }
    return nil
}
func insert_to_db(t_purchase_order* purchase_order,t *DeliverGoodsForPO)error {
		var err error
		 err=check_po_exist(t_purchase_order.po_no)
		 if err!=nil{
	    		return insert_goods_delivery_note(t_purchase_order,t)
		 }
	
    _, err = db.Exec(
        `INSERT INTO t_purchase_order(
	    purchase_order_id,po_no,po_date,status,company_id,vendor_basic_id,
		contact_account_id,payment_terms,requested_delivery_date,
		shipping_method_id,destination_country_id,loading_port,
		certificate,po_url,total_quantity,total_amount,currency_id,comments,
		note,createAt,createBy,dr,data_version) 
		VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,t_purchase_order.purchase_order_id,
			t_purchase_order.po_no,
			t_purchase_order.po_date,
			t_purchase_order.status,
			t_purchase_order.company_id,
			t_purchase_order.vendor_basic_id,
			t_purchase_order.contact_account_id,
			t_purchase_order.payment_terms,
			t_purchase_order.requested_delivery_date,
			t_purchase_order.shipping_method_id,
			t_purchase_order.destination_country_id,
			t_purchase_order.loading_port,
			t_purchase_order.certificate,
			t_purchase_order.po_url,
			t_purchase_order.total_quantity,
			t_purchase_order.total_amount,
			t_purchase_order.currency_id,
			t_purchase_order.comments,
			t_purchase_order.note,
			t_purchase_order.createAt,
			t_purchase_order.createBy,
		  	t_purchase_order.dr,
		  	t_purchase_order.data_version)
	    if err!=nil{
	    	return err
	    }else{
	    	err= insert_purchase_order_detail(t_purchase_order,t)
	    	if(err!=nil){
	    		return err
	    	}else{
	    		err=insert_goods_delivery_note(t_purchase_order,t)
	    	}
	    }
   return err
}