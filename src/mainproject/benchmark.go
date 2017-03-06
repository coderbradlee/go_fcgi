package main
 
import (
    "fmt"
    // "logger"
    "net/http"
    "bytes"
    "time"
    )
func client_req() {
    url := "http://172.18.100.85:9888/po/deliver_goods"

    var jsonStr = []byte(post_json)
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("X-Custom-Header", "myvalue")
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    // fmt.Println("response Headers:", resp.Header)
    // body, _ := ioutil.ReadAll(resp.Body)
    // fmt.Println("response Body:", string(body))
}
func benchmark() {
    start := time.Now()
    
    for i := 50; i > 0; i-- {
        go client_req()
    }
    // seconds := 10
    // fmt.Print(time.Duration(seconds)*time.Second) // prints 10s
    // const (
    //         Nanosecond  Duration = 1
    //         Microsecond          = 1000 * Nanosecond
    //         Millisecond          = 1000 * Microsecond
    //         Second               = 1000 * Millisecond
    //         Minute               = 60 * Second
    //         Hour                 = 60 * Minute
    // )
    fmt.Printf("%dms\n", time.Since(start).Nanoseconds()/1000/1000/50)
}

var post_json=`{    "operation": "DeliverGoodsForPO",    "data": {        "request_system": 1,        "request_time": "2017-02-16 08:00:00",        "purchase_order": {            "bill_type": "Purchase Order",            "po_no": "PO-FR-20170216-001014",            "po_url": "/root/go_fcgi/go_fcgi",            "po_date": "2017-02-16 18:00:00",            "created_by": "",            "approved_by": "",            "status": 1,            "supplier": "Renesola Shanghai",            "website": "France",            "company": "ReneSola France",            "requested_delivery_date": "2017-03-20 24:00:00",            "trade_term": "EXW",            "payment_terms": "",            "ship_via": "Sea",            "destination_country": "France",            "loading_port": "Amsterdam",            "certificate": "",            "total_quantity": 2400,            "total_amount": 5690.47,            "currency": "EUR",            "comments": "",            "note": "",            "detail": [                {                    "product_name": "Highbay",                    "product_code": "RHB120X0302",                    "item_no": "3518020400845",                    "unit_price": 3.64,                    "quantity": 1000,                    "uom": "PCS",                    "sub_total": 3640,                    "warranty": 3,                    "comments": "",                    "note": ""                },                {                    "product_name": "Flood Light",                    "product_code": "RFL400AK01D06",                    "item_no": "3518030601741",                    "unit_price": 6.89,                    "quantity": 200,                    "uom": "PCS",                    "sub_total": 1378,                    "warranty": 3,                    "comments": "",                    "note": ""                }            ]        },        "deliver_notes": [            {                "supplier": "Renesola Shanghai",                "buyer": "",                "loading_port": "Amsterdam",                "trade_term": "CIF",                "ship_via": "Sea",                "packing_method": "Pallet",                "logistic": "DHL",                "logistic_contact": "",                "logistic_contact_email": "",                "logistic_contact_telephone_number": "",                "etd": "2017-02-28 17:00:00",                "eta": "2017-03-17 10:00:00",                "customs_clearance_date": "2017-03-18 10:00:00",                "total_freight_charges": 879.65,                "total_insurance_fee": 262,                "total_excluded_tax": 3650.65,                "currency": "EUR",                "commercial_invoice": {                    "ci_no": "CI-FR-20170226-000196",                    "ci_url": "/opt/renesola/apollo/file/ci/CI-FR-20170226-000196.pdf",                    "ci_date": "2017-02-16 18:00:00",                    "status": 1,                    "company": "ReneSola France",                    "invoice_type": 0,                    "total_amount": 5690.47,                    "currency": "EUR",                    "created_by": "",                    "approved_by": "",                    "note": ""                },                "packing_list": {                    "pl_no": "PKL-FR-20170226-000196",                    "pl_url": "/opt/renesola/apollo/file/pkl/PKL-FR-20170226-000196.pdf"                },                "bill_of_lading": {                    "bl_no": "",                    "bl_url": ""                },                "associated_so": {                    "associated_so_no": "SC-FR-20170226-000196",                    "associated_so_url": "/opt/renesola/apollo/file/sc/SC-FR-20170226-000196.pdf"                },                "detail": [                    {                        "product_name": "Highbay",                        "product_code": "RHB120X0302",                        "item_no": "3518020400845",                        "unit_price": 3.64,                        "quantity": 500,                        "uom": "PCS",                        "sub_total": 1820                    },                    {                        "product_name": "Flood Light",                        "product_code": "RFL400AK01D06",                        "item_no": "3518030601741",                        "unit_price": 6.89,                        "quantity": 100,                        "uom": "PCS",                        "sub_total": 689                    }                ]            }        ]    }}`