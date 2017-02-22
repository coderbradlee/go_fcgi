 package main
 import (
    // "time"
    // "log"
)
const(
    error_json_decode="-100"
    error_json_encode="-101"
    error_db_insert="-102"
    error_check="-103"
)
func check_data(origi *DeliverGoodsForPO)(string,error) {

    return error_check,error.New("can't work with -103")
    
}
