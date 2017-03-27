package main
 import (
    "fmt"
    "time"
    // "errors"
    "net/http"
    "encoding/json"
    "io/ioutil"
    "strconv"
    "strings"
    "logger"
)
func get_purchase_order_id_chan(purchase_order_id_chan chan<- string,po_no string) {
    var purchase_order_id string
    db.QueryRow(fmt.Sprintf("select purchase_order_id from t_purchase_order where associated_po_no like '%%%s%%'",po_no)).Scan(&purchase_order_id)
     purchase_order_id_chan<-purchase_order_id
}
func get_payment_type_id_chan(payment_type_id_chan chan<- string,payment_type,company_id string) {
    var payment_type_id string
    db.QueryRow(fmt.Sprintf("select payment_type_id from t_payment_type where (short_name like '%%%s%%' or name like '%%%s%%') and company_id='%s' and status=0",payment_type,payment_type,company_id)).Scan(&payment_type_id)
    payment_type_id_chan<-payment_type_id
}
func get_payment_method_id_chan(payment_method_id_chan chan<- string,payment_method,company_id string) {
    var payment_method_id string
    db.QueryRow(fmt.Sprintf("select payment_method_id from t_payment_method where (short_name like '%%%s%%' or name like '%%%s%%') and company_id='%s' and status=0",payment_method,payment_method,company_id)).Scan(&payment_method_id)
    payment_method_id_chan<-payment_method_id
}
func get_payment_term_id_chan(payment_term_id_chan chan<- string,payment_term,company_id string) {
    s := strings.Split(payment_term, "|")
    payment_method:=s[0]
    payment_type:=s[1]
    
    payment_type_id_chan :=make(chan string)
    go get_payment_type_id_chan(payment_type_id_chan,payment_type,company_id)
    payment_type_id:=<-payment_type_id_chan

    payment_method_id_chan :=make(chan string)
    go get_payment_method_id_chan(payment_method_id_chan,payment_method,company_id)
    payment_method_id:=<-payment_method_id_chan
    
    fmt.Printf("payment_terms:%s,%s,%s",payment_method_id,payment_type_id,company_id)

    var payment_term_id string
    // payment_type_id payment_method_id
    db.QueryRow(fmt.Sprintf("select payment_term_id from t_payment_term where payment_type_id=? and payment_method_id=? and company_id=?",payment_type_id,payment_method_id,company_id)).Scan(&payment_term_id)
    if payment_term_id!=""{
        payment_term_id_chan<-payment_term_id
    }else{
        payment_term_id=rand_string(20)
        _, err := db.Exec(
        `INSERT INTO t_payment_term(payment_term_id,payment_type_id,payment_method_id) 
        VALUES (?,?,?)`,payment_term_id,payment_type_id,payment_method_id)
        if err!=nil{
            payment_term_id_chan<-""
            logger.Error(fmt.Sprintf("insert t_payment_term: %s", err.Error()))
        }else{
            payment_term_id_chan<-payment_term_id
        }
    }
}

func get_system_account_id_chan(system_account_id_chan chan<- string,name string) {
    var system_account_id string
    db.QueryRow(fmt.Sprintf("select system_account_id from t_system_account where account_name like '%%%s%%'",name)).Scan(&system_account_id)
     system_account_id_chan<-system_account_id
}
func get_bill_type_id_chan(bill_type_id_chan chan<- string,bill_type string) {
    var bill_type_id string
    db.QueryRow(fmt.Sprintf("select bill_type_id from t_bill_type where name like '%%%s%%'",bill_type)).Scan(&bill_type_id)
     bill_type_id_chan<-bill_type_id
}
func get_purchase_order_table_chan(purchase_order_table_chan chan<- purchase_order_part,po_no string) {
    var purchase_order_table purchase_order_part
    db.QueryRow(fmt.Sprintf("select purchase_order_id,company_id,vendor_basic_id,contact_account_id,payment_term_id,requested_delivery_date from t_purchase_order where associated_po_no='%s'",po_no)).Scan(
        &purchase_order_table.purchase_order_id,
        &purchase_order_table.company_id,
        &purchase_order_table.vendor_basic_id,
        &purchase_order_table.contact_account_id,
        &purchase_order_table.payment_term_id,
        &purchase_order_table.requested_delivery_date)
     purchase_order_table_chan<-purchase_order_table
}
func get_trade_term_id_chan(trade_term_id_chan chan<- string,Trade_term string) {
    var trade_term_id string
    db.QueryRow("select trade_term_id from t_trade_term where short_name=?",Trade_term).Scan(&trade_term_id)
    trade_term_id_chan<-trade_term_id
}
func get_buyer_id_chan(buyer_id_chan chan<- string,buyer string) {
    var buyer_id string
    db.QueryRow("select system_account_id from t_system_account where account_name=?",buyer).Scan(&buyer_id)
    buyer_id_chan<-buyer_id
}

func get_transport_term_id_chan(transport_term_id_chan chan<- string,ship_via string) {
    var transport_term_id string
    db.QueryRow("select ship_via_id from t_ship_via where full_name=?",ship_via).Scan(&transport_term_id)
    transport_term_id_chan<-transport_term_id
}
func get_packing_method_id_chan(packing_method_id_chan chan<- string,Packing_method string) {
    var packing_method_id string
    db.QueryRow("select packing_method_id from t_packing_method where name=?",Packing_method).Scan(&packing_method_id)
    packing_method_id_chan<- packing_method_id
}
func get_logistic_master_id_chan(logistic_master_id_chan chan<- string,Logistic string) {
    var logistic_master_id string
    db.QueryRow("select logistic_provider_master_id from t_logistic_provider_master where native_name=?",Logistic).Scan(&logistic_master_id)
    logistic_master_id_chan<-logistic_master_id
}
func get_logistic_contact_id_chan(logistic_contact_id_chan chan<- string,Logistic_contact string) {
    var logistic_contact_id string
    db.QueryRow("select logistic_provider_basic_id from t_logistic_provider_basic where contact=?",Logistic_contact).Scan(&logistic_contact_id)
    logistic_contact_id_chan<-logistic_contact_id
}
func get_currency_id(currency_id_chan chan<- string,currency string){
    var currency_id string
    db.QueryRow("select currency_id from t_currency where code=?",currency).Scan(&currency_id)
    currency_id_chan<-currency_id
}

func get_flow_no(company,bill_type string)(string,error) {
    // var flow_no string
    //GET http://172.18.100.85:8088/flowNo/2/FR/PO
    url:=configuration.Redis_url+"/"+configuration.System_no+"/"+company+"/"+bill_type

    resp, err1 := http.Get(url)
    if err1 != nil {
        return  "",err1
    }

    defer resp.Body.Close()
    body, err2 := ioutil.ReadAll(resp.Body)
    if err2 != nil {
        // handle error
        return  "",err2
    }
    var data flow_no_json
    json.Unmarshal(body, &data)
    i, err3 := strconv.Atoi(data.FlowNo)
    if err3 != nil {
        // handle error
        return  "",err3
    }
    // str := string.format(%06d",i)
    str := fmt.Sprintf("%06d",i)
    return str,nil
}
func get_subflow_no(company,parent_type,parent_no,bill_type string)(string,error) {
    // GET http://172.18.100.85:8088/subFlowNo/2/FR/PO/000196/GDN
    url:=configuration.Redis_url_subflowno+"/"+configuration.System_no+"/"+company+"/"+parent_type+"/"+parent_no+"/"+bill_type

    resp, err1 := http.Get(url)
    if err1 != nil {
        return  "",err1
    }

    defer resp.Body.Close()
    body, err2 := ioutil.ReadAll(resp.Body)
    if err2 != nil {
        // handle error
        return  "",err2
    }
    var data flow_no_json
    json.Unmarshal(body, &data)
    i, err3 := strconv.Atoi(data.FlowNo)
    if err3 != nil {
        // handle error
        return  "",err3
    }
    // str := string.format(%06d",i)
    str := fmt.Sprintf("%03d",i)
    return str,nil
}
// func get_goods_delivery_note_no(deliver_note_no string)(string,error) {
    // goods_delivery_note_no:="GDN-"
    // var short string
    // db.QueryRow("select note from t_company where short_name=?",company).Scan(&short)
    // // QU-UK-20160930-000001
    // goods_delivery_note_no+=short+"-"
    // goods_delivery_note_no+=time.Now().Format("20060102")+"-"
    // flow,err:=get_flow_no(short)
    // if err!=nil{
    //     return "",err
    // }
    // // flow:="000001"
    // goods_delivery_note_no+=flow//get int,format to 6bit,then convert to string
    // return goods_delivery_note_no,nil
    //check if deliver_note_no already exist in t_goods_delivery_note
// }
func get_vendor_master_id_chan(vendor_master_id_chan chan<- string,vendor_basic_id string) {
    var vendor_master_id string
    db.QueryRow("select vendor_master_id from t_vendor_master where vendor_basic_id=?",vendor_basic_id).Scan(&vendor_master_id)
    vendor_master_id_chan<-vendor_master_id
}
func get_company_short_name_chan(company_short_name_chan chan<- string,company string) {
    var company_short_name string
    db.QueryRow(fmt.Sprintf("select note from t_company where short_name like '%%%s%%'",company)).Scan(&company_short_name)
    company_short_name_chan<-company_short_name
}
func get_company_id_chan(company_id_chan chan<- string,company string) {
    var company_id string
    db.QueryRow(fmt.Sprintf("select company_id from t_company where short_name like '%%%s%%'",company)).Scan(&company_id)
    // db.QueryRow("select company_id from t_company where short_name like '%%?%%'",company).Scan(&company_id)
    company_id_chan<-company_id
}
func get_country_id_chan(country_id_chan chan<- string,country string) {
    var country_id string
    db.QueryRow(fmt.Sprintf("select country_id from t_country where short_name like '%%%s%%'",country)).Scan(&country_id)
    country_id_chan<-country_id
}
func get_port_id_chan(port_id_chan chan<- string,port string) {
    var port_id string
    db.QueryRow(fmt.Sprintf("select seaport_id from t_seaport where full_name like '%%%s%%'",port)).Scan(&port_id)
    port_id_chan<-port_id
}
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