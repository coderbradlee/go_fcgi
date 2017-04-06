 package main
 import (
	"errors"
	// "strings"
	// "errgroup"
	"time"
	"fmt"
)
// func get_company_id(company string) string{
// 	// var item_basic_id string
//  //    db.QueryRow("select item_basic_id from t_item_basic where item_no=?",item_no).Scan(&item_basic_id)

// 	var company_id string
//     db.QueryRow("select company_id from t_company where short_name=?",company).Scan(&company_id)

//     return company_id
// }

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

///先后顺序及优先级
//1、t_purchase_order 2、t_purchase_order_detail
//3、t_goods_delivery_note 3、t_commercial_invoice
//4、t_goods_delivery_note_detail 4、t_goods_delivery_note_attachment
//4、t_goods_receipt
func insert_to_db(t_purchase_order* purchase_order,t *PoData,sd *shared_data)(string,error) {
	// var level3_group errgroup
	// var level4_group errgroup
	var exist bool=false
		var err error
		var s string
		exist,err=check_po_exist(t_purchase_order)
		 
 		if exist{//err!=nil also does not exist
 			// level3_group.Go(t_purchase_order,t,sd,insert_goods_delivery_note)
		 	// level3_group.Go(t_purchase_order,t,sd,insert_commercial_invoice)
		 	// if s,err = level3_group.Wait(); err != nil {
		 	// 	return s,err
		 	// }else{
		 	// 	level4_group.Go(t_purchase_order,t,sd,insert_note_attachment)
    // 			level4_group.Go(t_purchase_order,t,sd,insert_note_detail)
    // 			level4_group.Go(t_purchase_order,t,sd,insert_goods_receipt)
    // 			s,err = level4_group.Wait()
    // 			return s,err
 			// }	
 			return error_check_po_exists,errors.New("po_no already exists")
 		}
 		company_short_name_chan :=make(chan string)
	    go get_company_short_name_chan(company_short_name_chan,t.Data.Purchase_order.Company)
	    company_short_name:=<-company_short_name_chan
	    /////////////////////////////////////
		payment_term_id_chan :=make(chan string)
	    go get_payment_term_id_chan(payment_term_id_chan,t_purchase_order.payment_terms,t_purchase_order.company_id)
	    t_purchase_order.payment_terms=<-payment_term_id_chan
	    //////////////////////////////////////////////////
		flow_no,err:=get_flow_no(company_short_name,"PO")
	    if flow_no==""{
	        return error_get_flow_no_po,errors.New("get_flow_no_po error")
	    }
	    po_no:="PO-"+company_short_name+"-"+time.Now().Format("20060102")+"-"+flow_no
	    fmt.Println("po_no:",po_no)
    _, err = db.Exec(
        `INSERT INTO t_purchase_order(
	    purchase_order_id,po_no,associated_po_no,po_date,status,company_id,vendor_basic_id,
		contact_account_id,payment_term_id,requested_delivery_date,
		shipping_method_id,destination_country_id,loading_port,unloading_port,po_url,total_quantity,total_amount,currency_id,comments,
		note,createAt,createBy,updateBy,dr,data_version) 
		VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,t_purchase_order.purchase_order_id,
			po_no,
			t_purchase_order.po_no,
			t_purchase_order.po_date,
			5,// t_purchase_order.status,
			t_purchase_order.company_id,
			t_purchase_order.vendor_basic_id,
			t_purchase_order.contact_account_id,
			t_purchase_order.payment_terms,
			t_purchase_order.requested_delivery_date,
			t_purchase_order.shipping_method_id,
			t_purchase_order.destination_country_id,
			t_purchase_order.loading_port,
			t_purchase_order.unloading_port,
			// t_purchase_order.certificate,
			t_purchase_order.po_url,
			t_purchase_order.total_quantity,
			t_purchase_order.total_amount,
			t_purchase_order.currency_id,
			t_purchase_order.comments,
			t_purchase_order.note,
			t_purchase_order.createAt,
			t_purchase_order.createBy,
			"go_fcgi",
		  	t_purchase_order.dr,
		  	t_purchase_order.data_version)
	    if err!=nil{
	    	return error_insert_purchase_order,err
	    }else{
	    	s,err= insert_purchase_order_detail(t_purchase_order,t,sd)
	    	if err!=nil{
	    		return s,err
	    	}else{
	   //  	level3_group.Go(t_purchase_order,t,sd,insert_goods_delivery_note)
		 	// level3_group.Go(t_purchase_order,t,sd,insert_commercial_invoice)
		 	// if s,err = level3_group.Wait(); err != nil {
		 	// 	return s,err
		 	// }else{
		 	// 	level4_group.Go(t_purchase_order,t,sd,insert_note_attachment)
    // 			level4_group.Go(t_purchase_order,t,sd,insert_note_detail)
    // 			level4_group.Go(t_purchase_order,t,sd,insert_goods_receipt)
    // 			s,err = level4_group.Wait()
    // 			return s,err
		 	// }
	   		return "",nil
		}
	}
}
func check_to_db(t_purchase_order* purchase_order,t *PoData,sd *shared_data)(string,error) {
	// var level3_group errgroup
	// var level4_group errgroup
	var exist bool=false
		var err error
		var s string
		exist,err=check_po_exist(t_purchase_order)
		 
 		if exist{//err!=nil also does not exist
 			// level3_group.Go(t_purchase_order,t,sd,insert_goods_delivery_note)
		 	// level3_group.Go(t_purchase_order,t,sd,insert_commercial_invoice)
		 	// if s,err = level3_group.Wait(); err != nil {
		 	// 	return s,err
		 	// }else{
		 	// 	level4_group.Go(t_purchase_order,t,sd,insert_note_attachment)
    // 			level4_group.Go(t_purchase_order,t,sd,insert_note_detail)
    // 			level4_group.Go(t_purchase_order,t,sd,insert_goods_receipt)
    // 			s,err = level4_group.Wait()
    // 			return s,err
 			// }	
 			return error_check_po_exists,errors.New("po_no already exists")
 		}
 		company_short_name_chan :=make(chan string)
	    go get_company_short_name_chan(company_short_name_chan,t.Data.Purchase_order.Company)
	    company_short_name:=<-company_short_name_chan
	    /////////////////////////////////////
		payment_term_id_chan :=make(chan string)
	    go get_payment_term_id_chan(payment_term_id_chan,t_purchase_order.payment_terms,t_purchase_order.company_id)
	    t_purchase_order.payment_terms=<-payment_term_id_chan
	    //////////////////////////////////////////////////
		flow_no,err:=get_flow_no(company_short_name,"PO")
	    if flow_no==""{
	        return error_get_flow_no_po,errors.New("get_flow_no_po error")
	    }
	    po_no:="PO-"+company_short_name+"-"+time.Now().Format("20060102")+"-"+flow_no
	    fmt.Println("po_no:",po_no)
    	err = nil
	    if err!=nil{
	    	return error_insert_purchase_order,err
	    }else{
	    	s,err= check_purchase_order_detail(t_purchase_order,t,sd)
	    	if err!=nil{
	    		return s,err
	    	}else{
	   		return "",nil
		}
	}
}