 package main
 import (
    // "time"
    "errors"
)
const(
    error_json_decode="-100"
    error_json_encode="-101"
    error_db_insert="-102"
    error_check_request_system="-120"
    error_check_bill_type="-121"
)
func check_request_system(request_system int32)error {
    if request_system!=1{
        return errors.New("request_system !=1")
    }
    return nil
}
func check_bill_type() {
    if bill_type!="Purchase Order"{
        return errors.New("bill_type!="Purchase Order"")
    }
    return nil
}
func check_data(origi *DeliverGoodsForPO)(string,error) {
    var err error
    err=check_request_system(origi.Data.Request_system)
    if err!=nil{
        return error_check_request_system,err
    }
    err=check_bill_type(origi.Data.Purchase_order.Bill_type)
    if err!=nil{
        return error_check_bill_type,err
    }
    return "",nil
}
