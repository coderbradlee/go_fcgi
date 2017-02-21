 package main
 import (
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
    db.QueryRow("select shipping_via_id from t_ship_via where full_name=?",Ship_via).Scan(&shipping_via_id)
    return shipping_via_id
}
