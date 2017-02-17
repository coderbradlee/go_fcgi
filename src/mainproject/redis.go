 package main
 
import (
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "math/rand"
    // "crypto/rand"
    "time"
    // "log"
    // "math/big"
)
/**
 * copy inventory_balance data to cost_statistics struct
 */
//  func rand_string(length int) string {
//     token := ""
//     codeAlphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
//     codeAlphabet += "0123456789"

//     for i := 0; i < length; i++ {
//         token += string(codeAlphabet[cryptoRandSecure(int64(len(codeAlphabet)))])
//     }
//     return token
// }

// func cryptoRandSecure(max int64) int64 {
//     // fmt.Println(time.Now().UnixNano())
//     rand.Seed(time.Now().UnixNano())
//     nBig, err := rand.Int(rand.Reader, big.NewInt(max))
//     if err != nil {
//         log.Println(err)
//     }
//     return nBig.Int64()
// }
func rand_string(lens int)string{
    choice:="ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789";
    var ret string
    rand.Seed(time.Now().UnixNano())
    for i:=0;i<lens;i++{
        ret+=string(choice[rand.Intn(35)]);
    }
   return ret;
}
type cost_statistics struct {
    cost_statistics_id string
 	company_id string
 	product_category_id string
    item_master_id string
    total_stock_quantity int32
 	stock_uom_id string
 	total_cost_amount int32
    last_month_unit_price int32
 	currency_id string
    statistic_beginning_date string
    statistic_ending_date string
    accounting_period_id string
    sort_no int32
    createAt string
    dr int32
    data_version int32
}
func (u *cost_statistics)print() {
    fmt.Printf("%s|%s|%s|%s|%d|%s|%d|%d|%s|%s|%s|%s|%d|%s|%d|%d\n", u.cost_statistics_id,u.company_id,u.product_category_id,u.item_master_id,u.total_stock_quantity,u.stock_uom_id,u.total_cost_amount,u.last_month_unit_price,u.currency_id,u.statistic_beginning_date,u.statistic_ending_date,u.accounting_period_id,u.sort_no,u.createAt,u.dr,u.data_version)
}
var g_insert_data []*cost_statistics
func copy(records []*inventory_balance) {
    for _,i:=range records {
        p := new(cost_statistics)
        p.company_id=i.company_id
        p.accounting_period_id=i.accounting_period_id
        p.item_master_id=i.item_master_id
        p.total_stock_quantity=i.in_stock_balance
        p.stock_uom_id=i.uom_id
        g_insert_data = append(g_insert_data, p)
    }
}
func pad_product_category_id(g_insert_data []*cost_statistics) {
    for _,i:=range g_insert_data {
        company_id:=i.company_id
        item_master_id:=i.item_master_id
        i.product_category_id=get_product_category_id(company_id,item_master_id)
        i.dr=0
        i.data_version=1
        i.cost_statistics_id=rand_string(20)
        // i.print()
    }
}
func pad_statistic_time(g_insert_data []*cost_statistics) {
    for _,i:=range g_insert_data {
        i.statistic_beginning_date,i.statistic_ending_date=get_statistic_time(i.accounting_period_id)
        // i.print()
    }
}
func pad_createAt(g_insert_data []*cost_statistics) {
    for _,i:=range g_insert_data {
        t := time.Now()
        i.createAt=t.Format("2006-01-02 15:04:05")
        // i.print()
    }
}
func pad_currency_id(g_insert_data []*cost_statistics) {
    for _,i:=range g_insert_data {
        i.currency_id=get_currency_id(i.company_id,i.item_master_id)
        i.print()
    }
}
func get_currency_id(company_id string,item_master_id string)string {
    var item_basic_id string
    err := db.QueryRow("select item_basic_id from t_item_master where item_master_id=? and company_id=?",item_master_id,company_id).Scan(&item_basic_id)

    var currency_id string
    err = db.QueryRow("select currency_id from t_sales_price_book where item_basic_id=? and company_id=?",item_basic_id,company_id).Scan(&currency_id)
    checkErr(err)
    return currency_id
}
func get_statistic_time(accounting_period_id string)(string,string) {
    var statistic_beginning_date string
    var statistic_ending_date string
    err := db.QueryRow("select opening_date,ending_date from t_accounting_period where accounting_period_id=?",accounting_period_id).Scan(&statistic_beginning_date,&statistic_ending_date)
    checkErr(err)
    return statistic_beginning_date,statistic_ending_date
}
func get_product_category_id(company_id string,item_master_id string)string {
    
    var item_basic_id string
    err := db.QueryRow("select item_basic_id from t_item_master where item_master_id=? and company_id=?",item_master_id,company_id).Scan(&item_basic_id)
    
    var item_category_id string
    err = db.QueryRow("select item_category_id from t_item_basic where item_basic_id=?",item_basic_id).Scan(&item_category_id)
    
    var product_category_id string
    err = db.QueryRow("select product_category_id from t_product_category_item_category_link where item_category_id=?",item_category_id).Scan(&product_category_id)
    
    checkErr(err)
    return product_category_id
}