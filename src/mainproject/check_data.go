 package main
 import (
    // "time"
    "errors"
    "os"
    // "fmt"
    // "reflect"
)
const(
    error_json_decode="-100"//json 解包错误
    error_json_encode="-101"//json 打包错误
    error_db="-102"//连接mysql数据库错误
    error_check_request_system="-120"//请求系统是否为1
    error_check_bill_type="-121"//bill_tpye是否为Purchase Order
    error_check_po_no="-122"//po_no长度问题，数据库表里面是20位
    error_check_po_url="-123"//是否存在文件
    error_check_status="-124"//status是否为1
    error_check_supplier="-125"//supplier是否为Renesola Shanghai
    error_check_packing_method="-126"//t_packing_method表里没有此packing_method
    error_check_logistic_provider="-127"//物流提供商信息缺失
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
func check_supplier(supplier string)error {
    if supplier!="Renesola Shanghai"{
        return errors.New(`supplier is not Renesola Shanghai`)
    }
    return nil
}
func check_packing_method(deliver_notes []Deliver_notes)error {
    // fmt.Println(deliver_notes[0].Packing_method)
    for _,d:=range deliver_notes{
        // fmt.Println(reflect.TypeOf(d))
        var packing_method string
        db.QueryRow("select packing_method_id from t_packing_method where native_name=?",d.Packing_method).Scan(&packing_method)
        if packing_method== ""{
            return errors.New(`packing_method_id missed`)
        }
    }
    return nil
}
func check_logistic_provider(deliver_notes []Deliver_notes)error {
    // fmt.Println(deliver_notes[0].Packing_method)
    for _,d:=range deliver_notes{
        // fmt.Println(reflect.TypeOf(d))
        var logistic_provider_basic_id string
        db.QueryRow("select logistic_provider_basic_id from t_logistic_provider_basic where name=?",d.Logistic).Scan(&logistic_provider_basic_id)
        if logistic_provider_basic_id== ""{
            return errors.New(`logistic_provider_basic_id missed`)
        }
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
    err=check_supplier(origi.Data.Purchase_order.Supplier)
    if err!=nil{
        return error_check_supplier,err
    }
    err=check_packing_method(origi.Data.Deliver_notes)
    if err!=nil{
        return error_check_packing_method,err
    }
    err=check_logistic_provider(origi.Data.Deliver_notes)
    if err!=nil{
        return error_check_logistic_provider,err
    }
    
    return "",nil
}
