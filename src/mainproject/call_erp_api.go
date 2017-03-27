 package main
 import (
    "time"
    "logger"
    "fmt"
    "net/http"
    "io/ioutil"
    // "strconv"
    "encoding/json"
    "errors"
    // "bytes"
    "strings"
)

func call_erp_api(gdn_nos []erp_api_data)(string,error) {
    var ret erp_api_return_json
    ret.Operation="SubmitGoodsDeliveryNote"
    ret.Request_time=time.Now().Format("2006-01-02 15:04:05")
    ret.Data.Action_name="DeliverGoods"
    for _,gdn_no:=range gdn_nos{
        // fmt.Printf("%s:%s:%s\n",gdn_no.company_id,gdn_no.goods_delivery_note_id,gdn_no.goods_delivery_note_no)
        ret.Data.Company_id=gdn_no.company_id
        var e erp_api_return_json_goods_delivery_notes
        e.Goods_delivery_note_id=gdn_no.goods_delivery_note_id
        e.Goods_delivery_note_no=gdn_no.goods_delivery_note_no
        e.Goods_delivery_note_status=0
        ret.Data.Goods_delivery_notes=append(ret.Data.Goods_delivery_notes,e)
    }
   
    // configuration.Erp_api
    // ret.Data=ret_data
    fmt.Printf("len:%d\n",len(ret.Data.Goods_delivery_notes))
    var b []byte
    if b, err := json.Marshal(ret); err == nil {
        fmt.Println(string(b))
    }else{
        logger.Error("json Marshal")
        return error_call_erp_api,errors.New("error_call_erp_api json Marshal")
    }

    return post_api(string(b))
}
func post_api(content string)(string,error) {
    ///////////////////post
    resp, err := http.Post(configuration.Erp_api,
        "application/x-www-form-urlencoded",strings.NewReader(content))
    
    if err != nil {
        fmt.Println(err)
        logger.Error(fmt.Sprintf("post %s :%s", configuration.Erp_api, content))

        return error_call_erp_api,err
    }
    fmt.Println("555555555555")
    body, err := ioutil.ReadAll(resp.Body)
    defer resp.Body.Close()
    if err != nil {
        // handle error
        logger.Error(fmt.Sprintf("reponse from %s :%s", configuration.Erp_api, string(body)))
        return error_call_erp_api,err
    }
    logger.Info(fmt.Sprintf("response from %s :%s", configuration.Erp_api, string(body)))
    fmt.Println(string(body))
    var t erp_api_reponse  
        // var po_shared_data shared_data
    err_decode := json.Unmarshal(body, &t)
    if err_decode!=nil{
        logger.Error(fmt.Sprintf("json unmarshal reponse from %s :%s", configuration.Erp_api, string(body)))
        return error_call_erp_api,err_decode
    }
    if t.Error_code!="200"{
        logger.Error(fmt.Sprintf("reponse !=200"))
        return error_call_erp_api,errors.New("reponse !=200 from erp_api")
    }
    return "",nil
}
type erp_api_reponse struct{
    Error_code string `json:"error_code"`
    Error_msg string `json:"error_msg"`
    Data erp_api_reponse_data `json:"data"`
    Reply_time string `json:"reply_time"`
}
type erp_api_reponse_data struct{
    Goods_delivery_notes []erp_api_reponse_data_gdns `json:"goods_delivery_notes"`
}
type erp_api_reponse_data_gdns struct{
    Goods_delivery_note_id string `json:"goods_delivery_note_id"`
    Goods_delivery_note_no string `json:"goods_delivery_note_no"`
    Wf_instance_id string `json:"wf_instance_id"`
}
type erp_api_data struct{
    company_id string
    goods_delivery_note_id string
    goods_delivery_note_no string
}
type erp_api_return_json_goods_delivery_notes struct{
    Goods_delivery_note_id string `json:"goods_delivery_note_id"`
    Goods_delivery_note_no string `json:"goods_delivery_note_no"`
    Goods_delivery_note_status int `json:"goods_delivery_note_status"`
}
type erp_api_return_json_data struct{
    Action_name string `json:"action_name"`
    Company_id string `json:"company_id"`
    Goods_delivery_notes []erp_api_return_json_goods_delivery_notes `json:"goods_delivery_notes"`       
}
type erp_api_return_json struct{
    Operation string `json:"operation"`
    Data erp_api_return_json_data `json:"data"`
    Request_time string `json:"request_time"`
}