 package main
 import (
    // "time"
    "errors"
    "os"
)
const(
    error_json_decode="-100"
    error_json_encode="-101"
    error_db_insert="-102"
    error_check_request_system="-120"
    error_check_bill_type="-121"
    error_check_po_no="-122"
    error_check_po_url="-123"
    error_check_status="-124"
)
func check_request_system(request_system int32)error {
    if request_system!=1{
        return errors.New("request_system !=1")
    }
    return nil
}
func check_bill_type(bill_type string)error {
    if bill_type!="Purchase Order"{
        return errors.New(`bill_type!=Purchase Order`)
    }
    return nil
}
func check_po_no(po_no string)error {
    cs:="PO-FR-20170216-00101"
    if len(po_no)>len(cs){
        return errors.New(`po_no is too long`)
    }
    return nil
}
func check_po_url(po_url string)error {
    _,err:=os.Stat(po_url)
    return err
}
func check_status(status int32)error {
    if status!=1{
        return errors.New(`status!=1`)
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
    err=check_po_no(origi.Data.Purchase_order.Po_no)
    if err!=nil{
        return error_check_po_no,err
    }
    err=check_po_url(origi.Data.Purchase_order.Po_url)
    if err!=nil{
        return error_check_po_url,err
    }
    err=check_status(origi.Data.Purchase_order.Status)
    if err!=nil{
        return error_check_status,err
    }
    return "",nil
}
