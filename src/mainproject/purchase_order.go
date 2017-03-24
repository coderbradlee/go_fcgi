 package main
 import (
	"errors"
	// "strings"
	// "errgroup"
	"time"
	"fmt"
)

func insert_to_db(t_purchase_order* purchase_order,t *PoData,sd *shared_data)(string,error) {
	// var level3_group errgroup
	// var level4_group errgroup
	var exist bool=false
		var err error
		var s string
		exist,err=check_po_exist(t_purchase_order)
		 
 		if exist{
 			return error_check_po_exists,errors.New("po_no already exists")
 		}
 		company_short_name_chan :=make(chan string)
	    go get_company_short_name_chan(company_short_name_chan,t.Data.Purchase_order.Company)
	    company_short_name:=<-company_short_name_chan
	    //////////////////////////////////////////////////
		flow_no,err:=get_flow_no(company_short_name,"PO")
	    if flow_no==""{
	        return error_get_flow_no_po,errors.New("get_flow_no_po error")
	    }
	    po_no:="PO-"+company_short_name+"-"+time.Now().Format("20060102")+"-"+flow_no
	    fmt.Println("po_no:",po_no)
	    stmt, err := db.Prepare(
        `INSERT INTO t_purchase_order(
	    purchase_order_id,po_no,associated_po_no,po_date,status,company_id,vendor_basic_id,
		contact_account_id,payment_terms,requested_delivery_date,
		shipping_method_id,destination_country_id,loading_port,unloading_port,
		certificate,po_url,total_quantity,total_amount,currency_id,comments,
		note,createAt,createBy,updateBy,dr,data_version) 
		VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`)
		if err != nil {
		    return error_insert_purchase_order,err
		}
		_, err = stmt.Exec(t_purchase_order.purchase_order_id,
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
			t_purchase_order.certificate,
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
		if err != nil {
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