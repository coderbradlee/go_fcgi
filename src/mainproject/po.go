 package main
 import (
    "log"
    "fmt"
    _"encoding/json"
    "net/http"
    "io/ioutil"
)
type Purchase_order struct{
 bill_type string
//          "po_no":"PO-FR-20170216-001014",
//          "po_url":"/opt/renesola/apollo/file/po/PO-FR-20170216-001014.pdf",
//          "po_date":"2017-02-16 18:00:00",
//          "create_by":"",
//          "status":1,
// "supplier":"Renesola Shanghai",
//          "website":"France",
//          "company":"ReneSola France",
//          "requested_delivery_date":"2017-03-20 24:00:00",
// "trade_term":"EXW",
//          "payment_terms":"",
//          "ship_via":"Sea",
//          "destination_country":"France",
//          "loading_port":"Amsterdam",
//          "certificate":"",
//          "total_quantity":2400,
//          "total_amount":5690.47,
//          "currency":"EUR",
//          "comments":"",
//          "note":"",
//          "detail":[{
//             "product_name":"Highbay",
//             "product_code":"RHB120X0302",
//             "item_no":"3518020400845",
//             "unit_price":3.64,
//             "quantity":1000,
//             "uom":"PCS",
//             "sub_total":3640.00,
//             "warranty":3,
//             "comments":"",
// "note":""
// },{
//             "product_name":"Flood Light",
//             "product_code":"RFL400AK01D06",
//             "item_no":"3518030601741",
//             "unit_price":6.89,
//             "quantity":200,
//             "uom":"PCS",
//             "sub_total":1378.00,
//             "warranty":3,
//             "comments":"",
// "note":""
// }]
// },
// "deliver_notes":[{
// "supplier":"Renesola Shanghai",
// " buyer":"",
// "loading_port":"Amsterdam",
// "trade_term":"CIF",
// "ship_via":"Sea",
// "packing_method":"Pallet",
// "logistic":"COSCO",
// "logistic_contact":"",
// "logistic_contact_email":"",
// "logistic_contact_telephone_number":"",
// "etd":"2017-02-28 17:00:00",
// "eta":"2017-03-17 10:00:00",
// "customs_clearance_date":"2017-03-18 10:00:00",
// "total_freight_charges":879.65,
// "total_insurance_fee":262.00,
// "total_excluded_tax":3650.65,
// "currency":"EUR",
// "commercial_invoice":{
// "ci_no":"CI-FR-20170226-000196",
// "ci_url":"/opt/renesola/apollo/file/ci/CI-FR-20170226-000196.pdf"
// },
// "packing_list":{
// "pl_no":"PKL-FR-20170226-000196",
// "pl_url":"/opt/renesola/apollo/file/pkl/PKL-FR-20170226-000196.pdf"
// },
// "bill_of_lading":{
// "bl_no":"",
// "bl_url":""
// },
// "associated_so":{
// "associated_so_no":"SC-FR-20170226-000196",
// "associated_so_url":"/opt/renesola/apollo/file/sc/SC-FR-20170226-000196.pdf"
// },
// "detail":[{
//             "product_name":"Highbay",
//             "product_code":"RHB120X0302",
//             "item_no":"3518020400845",
//             "unit_price":3.64,
//             "quantity":500,
//             "uom":"PCS",
//             "sub_total":1820.00
// },{
//             "product_name":"Flood Light",
//             "product_code":"RFL400AK01D06",
//             "item_no":"3518030601741",
//             "unit_price":6.89,
//             "quantity":100,
//             "uom":"PCS",
//             "sub_total":689.00
// }]
// }]
//    }
// }

// }
}
type Data struct{
	request_system int
	request_time string
	purchase_order Purchase_order
}
type DeliverGoodsForPO struct {
   operation string
   data Data 
}

func poHandler (w http.ResponseWriter, r *http.Request) {
	////////////////////////////////
	addr := r.Header.Get("X-Real-IP")
	if addr == "" {
		addr = r.Header.Get("X-Forwarded-For")
		if addr == "" {
			addr = r.RemoteAddr
		}
	}

	// log.Printf("Started %s %s for %s", r.Method, r.URL.Path, addr)

/////////////////////////////////////////////////////////////////
	if r.Method !="POST"{
		fmt.Fprint(w, "this interface should be post!")
	} else{
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
	 		log.Println("ioutil.ReadAll error", err) 
 		}
 		sbody :=string(body)
		// log.Println(sbody)
		log.Printf("Started %s %s for %s:%s", r.Method, r.URL.Path, addr,sbody)
		decoder := json.NewDecoder(sbody)
	    var t DeliverGoodsForPO   
	    err := decoder.Decode(&t)
	    if err != nil {
	        panic(err)
	    }
	    defer r.Body.Close()
	    // log.Println(t.Test)
	    ret := `{
				   "error_code":"200",
				   "error_msg":"Goods received successfully at 2017-03-17 12:00:00",
				   "data":{
				      "goods_receipt_no":"GR-FR-20170226-000196",
				      "bill_type":"Goods Receipt",
				      "receive_by":"Enie Yang",
				      "company":"ReneSola France",
				      "receive_at":"2017-03-17 12:00:00"
				   },
				   "reply_time":"2017-03-17 12:00:00"
			   }`
		// log.Logger(ret)
		fmt.Fprint(w, ret)
	}

} 
