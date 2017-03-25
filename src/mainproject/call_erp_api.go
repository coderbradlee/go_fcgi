 package main
 import (
    "time"
    "logger"
    "fmt"
    // "net/http"
    // "io/ioutil"
    // "strconv"
    "encoding/json"
    "errors"
    // "bytes"
)

func call_erp_api(gdn_nos []erp_api_data)(string,error) {
    var ret erp_api_return_json
    ret.Operation="SubmitGoodsDeliveryNote"
    ret.Request_time=time.Now().Format("2006-01-02 15:04:05")
    var ret_data erp_api_return_json_data
    ret_data.Action_name="DeliverGoods"
    // var ee []erp_api_return_json_goods_delivery_notes
    // var ee=make([]erp_api_return_json_goods_delivery_notes)
    for _,gdn_no:=range gdn_nos{
        fmt.Printf("%s:%s:%s\n",gdn_no.company_id,gdn_no.goods_delivery_note_id,gdn_no.goods_delivery_note_no)
        ret_data.Company_id=gdn_no.company_id
        var e erp_api_return_json_goods_delivery_notes
        e.Goods_delivery_note_id=gdn_no.goods_delivery_note_id
        e.Goods_delivery_note_no=gdn_no.goods_delivery_note_no
        e.Goods_delivery_note_status=0
        ret_data.Goods_delivery_notes=append(ret_data.Goods_delivery_notes,e)
        // ee=append(ee,e)
    }
    // fmt.Printf("len:%d",len(ee))
    // ret_data.Goods_delivery_notes=ee

    // configuration.Erp_api
    ret.Data=ret_data
    fmt.Printf("len:%d\n",len(ret.Data.Goods_delivery_notes))
    // var buffer bytes.Buffer
    // enc := json.NewEncoder(&buffer)
    // err_encode := enc.Encode(ret)
    var b []byte
    if b, err := json.Marshal(ret); err == nil {
        fmt.Println(string(b))
    }else{
        fmt.Println("================struct 到json str==")
        logger.Error("json Marshal")
        return error_call_erp_api,errors.New("error_call_erp_api json Marshal")
    }

    return post_api(string(b))
}
func post_api(content string)(string,error) {
    ///////////////////post
    // resp, err := http.Post("configuration.Erp_api",
    //     "application/x-www-form-urlencoded",
    //     strings.NewReader(buffer.String()))
    // if err != nil {
    //     // fmt.Println(err)
    //     return error_call_erp_api,err
    // }
 
    // defer resp.Body.Close()
    // body, err := ioutil.ReadAll(resp.Body)
    // if err != nil {
    //     // handle error
    //     return error_call_erp_api,err
    // }
 
    // fmt.Println(string(body))
    return "",nil
}
type erp_api_data struct{
    company_id string
    goods_delivery_note_id string
    goods_delivery_note_no string
}
type erp_api_return_json_goods_delivery_notes struct{
    Goods_delivery_note_id string `json:"goods_delivery_note_id"`
    Goods_delivery_note_no string `json:"goods_delivery_note_id"`
    Goods_delivery_note_status int `json:"goods_delivery_note_id"`
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