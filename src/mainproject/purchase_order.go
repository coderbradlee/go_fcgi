 package main
 import (
	// "errors"
	// "strings"
	"fmt"
	// "runtime"
)
// func get_company_id(company string) string{
// 	// var item_basic_id string
//  //    db.QueryRow("select item_basic_id from t_item_basic where item_no=?",item_no).Scan(&item_basic_id)

// 	var company_id string
//     db.QueryRow("select company_id from t_company where short_name=?",company).Scan(&company_id)

//     return company_id
// }
func get_shipping_method_id_chan(shipping_method_id_chan chan<- string,Ship_via string) {
	//cannot find the way to shipping_method_id
	// var shipping_method_id string
 //    db.QueryRow("select shipping_method_id from t_company where short_name=?",company).Scan(&shipping_method_id)
	var shipping_via_id string
    db.QueryRow("select ship_via_id from t_ship_via where full_name=?",Ship_via).Scan(&shipping_via_id)
    shipping_method_id_chan<- shipping_via_id
}
func get_vendor_basic_id_chan(vendor_basic_id_chan chan<- string,supplier string) {
	var vendor_basic_id string
    db.QueryRow("select vendor_basic_id from t_vendor_basic where short_name=?",supplier).Scan(&vendor_basic_id)
    vendor_basic_id_chan<-vendor_basic_id
}
// func get_contact_account_id(company_id string)string {
// 	var contact_account_id string//来自采购主动发起方公司的运营经理
//     db.QueryRow(`select  
// 		c.system_account_id
// 		from  
// 		(select *  from t_wf_role_def
// 		where dr=0
// 		and alias='Operation Manager'
// 		) a
// 		inner join 
// 		(select  *  from t_wf_role_resolve
// 		where dr=0
// 		and master_file_obj_id=?
// 		) b
// 		on a.wf_role_def_id=b.wf_role_def_id
// 		inner join  (select *  from t_system_account where dr=0) c
// 		on b.employee_id=c.employee_no
// 		order by a.alias`,company_id).Scan(&contact_account_id)
//     return contact_account_id
// }
func check_po_exist(po_no string)(int,error) {
	var get_po_no string
	var err error
    err =db.QueryRow("select po_no from t_purchase_order where po_no=?",po_no).Scan(&get_po_no)
    if err!=nil{
    	return 0,err
    }
    if get_po_no!=""{
    	return 1,nil//存在po_no
    }
    return 0,nil
}
//1、t_purchase_order 2、t_purchase_order_detail
//3、t_goods_delivery_note 3、t_commercial_invoice
//4、t_goods_delivery_note_detail 4、t_goods_delivery_note_attachment
//4、t_goods_receipt
func level3(level12_chan chan<- error,t_purchase_order* purchase_order,t *DeliverGoodsForPO,sd *shared_data) {
	var level3_chan=make(chan error) 
	go insert_goods_delivery_note(level3_chan,t_purchase_order,t,sd)
	go insert_commercial_invoice(level3_chan,t_purchase_order,t,sd)
	
	fmt.Println("purchase_order.go 71")
	for i:=0;i<2;i++{
		fmt.Println("purchase_order.go 73")
		t:=<-level3_chan
		fmt.Println("purchase_order.go 75")
		if t!=nil{
			level12_chan<- t
		}
	}
	fmt.Println("purchase_order.go 78")
	var level4_chan=make(chan error) 
	go level4(level4_chan,t_purchase_order,t,sd)
	temp:=<-level4_chan
	if temp!=nil{
		level12_chan<- temp
	}else{
		level12_chan<-nil
	}
	
}
func level4(level3_chan chan<- error,t_purchase_order* purchase_order,t *DeliverGoodsForPO,sd *shared_data) {
	var level4_chan=make(chan error)
	go insert_note_attachment(level4_chan,t_purchase_order,t,sd)
    go insert_note_detail(level4_chan,t_purchase_order,t,sd)   
	go insert_goods_receipt(level4_chan,t_purchase_order,t,sd)

	for i:=0;i<3;i++{
		t:=<-level4_chan
		if t!=nil{
			level3_chan<- t
		}
	}
	level3_chan<-nil
}
func insert_to_db(t_purchase_order* purchase_order,t *DeliverGoodsForPO,sd *shared_data)error {
		var err error
		var exist int
		var level3_chan=make(chan error) 
		fmt.Println("purchase_order.go:109")
		 exist,err=check_po_exist(t_purchase_order.po_no)
		 fmt.Println("purchase_order.go:111")
		 if err!=nil{//存在po_no
		 	fmt.Println("purchase_order.go:113")
		 	// return err
		 }else{
		 	if exist==1{
		 		fmt.Println("exist")
		 		level3(level3_chan,t_purchase_order,t,sd)
		 		t:=<-level3_chan
		 		return t
		 	}
		 }
	fmt.Println("purchase_order.go:120")
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
	    	fmt.Println("purchase_order.go:152")
	    	return err
	    }else{
	    	fmt.Println("purchase_order.go:155")
	    	err= insert_purchase_order_detail(t_purchase_order,t,sd)
	    	if(err!=nil){
	    		return err
	    	}else{
	    		level3(level3_chan,t_purchase_order,t,sd)
		 		err= <-level3_chan
	    	}
	    }
   return err
}