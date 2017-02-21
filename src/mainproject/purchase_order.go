 package main
 import (
    
    "time"
)
func get_company_id(company string) string{
	// var item_basic_id string
 //    db.QueryRow("select item_basic_id from t_item_basic where item_no=?",item_no).Scan(&item_basic_id)

	var company_id string
    db.QueryRow("select company_id from t_company where short_name=?",company).Scan(&company_id)

    return company_id
}
