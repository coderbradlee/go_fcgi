package main
 import (
    "fmt"
    "time"
    // "errors"
)
func get_company_time_zone_chan(company_time_zone_chan chan<- float64,company string) {
    var company_time_zone float64
    db.QueryRow("select time_zone from t_company where short_name=?",company).Scan(&company_time_zone)
     company_time_zone_chan<-company_time_zone
 }
 func set_company_time_zone(company string,sd *shared_data){
 	company_time_zone_chan :=make(chan float64)
    go get_company_time_zone_chan(company_time_zone_chan,company)
    t:=<-company_time_zone_chan
    t-=8
    sd.company_time_zone,_=time.ParseDuration(fmt.Sprintf("%fh",t))
	// fmt.Println(company_time_zone)
	//  ParseDuration
 }
 func get_contact_account_id_sh_chan(contact_account_id_sh_chan chan<- string,company string) {
	var company_id string//来自采购主动发起方公司的运营经理
    db.QueryRow(`select company_id from t_company where short_name=?`,company).Scan(&company_id)

	var contact_account_id string//来自采购主动发起方公司的运营经理
    db.QueryRow(`select  
		c.system_account_id
		from  
		(select *  from t_wf_role_def
		where dr=0
		and alias='Operation Manager'
		) a
		inner join 
		(select  *  from t_wf_role_resolve
		where dr=0
		and master_file_obj_id=?
		) b
		on a.wf_role_def_id=b.wf_role_def_id
		inner join  (select *  from t_system_account where dr=0) c
		on b.employee_id=c.employee_no
		order by a.alias`,company_id).Scan(&contact_account_id)
    contact_account_id_sh_chan<- contact_account_id
}
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
func get_item_master_id_chan(item_master_id_chan chan<- string,item_no,product_name,product_code string){
	var item_basic_id string
    db.QueryRow("select item_basic_id from t_item_basic where item_no=?",item_no).Scan(&item_basic_id)

	var item_master_id string
    db.QueryRow("select item_master_id from t_item_master where item_basic_id=? and product_code=? and product_name=?",item_basic_id,product_code,product_name).Scan(&item_master_id)

    item_master_id_chan<-item_master_id
}
func get_uom_id_chan(uom_id_chan chan<- string,uom string){
	var uom_id string
    db.QueryRow("select uom_id from t_uom where name=?",uom).Scan(&uom_id)
    uom_id_chan<-uom_id
}