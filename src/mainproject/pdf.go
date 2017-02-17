 package main
 
import (
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "net/http"
    "log"
    "time"
    "sort"
)
/**
 * 初步获取t_inventory_balance里面的数据
 */
type credit_black_list struct {
      credit_black_list_id string
      company_id string
      customer_master_id string
      status int
      total_overdue_days int
      total_overdue_amount float32
      currency_id string
      statistic_beginning_date string
      statistic_ending_date string
      accounting_period_id string
      sort_no int
      createAt string
      createBy string
      dr int
      data_version int
      sales_order_id string //template
      payment_dead_line string
}
type type_credit_black_list []*credit_black_list
var g_credit_black_list type_credit_black_list
//实现sort的三个接口
func (c type_credit_black_list) Len() int {
    return len(c)
}
func (c type_credit_black_list) Swap(i, j int) {
    c[i], c[j] = c[j], c[i]
}
func (c type_credit_black_list) Less(i, j int) bool {
    if(c[i].total_overdue_days>c[j].total_overdue_days){
        return true
    }else if(c[i].total_overdue_days==c[j].total_overdue_days){
        if(c[i].total_overdue_amount>c[j].total_overdue_amount){
            return true
        }
    }
    return false
}
type sales_order struct {
      customer_master_id string
      currency_id string
}
type commercial_invoice struct {
       company_id string
       invoice_no string
       invoice_date string
       sales_order_id string
       payment_dead_line string
       payment_due float32
       status int
}

func (u *commercial_invoice) print(){
    fmt.Printf("%s||%s||%s||%s||%s||%f\n", u.company_id,u.invoice_no,u.invoice_date,u.sales_order_id,u.payment_dead_line,u.payment_due)
}
func (u *credit_black_list) print(){
    fmt.Printf("%s||%s||%s||%d||%f||%s||%s||%s||%s||%s\n", u.credit_black_list_id,u.company_id,u.customer_master_id,u.total_overdue_days,u.total_overdue_amount,u.currency_id,u.sales_order_id,u.statistic_beginning_date,u.statistic_ending_date,u.accounting_period_id)
    // credit_black_list_id string
    //   company_id string
    //   customer_master_id string
    //   status int
    //   total_overdue_days int
    //   total_overdue_amount float32
    //   currency_id string
    //   statistic_beginning_date string
    //   statistic_ending_date string
    //   accounting_period_id string
    //   sort_no int
    //   createAt string
    //   createBy string
    //   dr int
    //   data_version int
    //   sales_order_id string //template
}
func credit_start(w http.ResponseWriter, r *http.Request) {
    //首先获取公司id，让后按公司id进行下面的工作
    rows, err := db.Query(`SELECT DISTINCT
       company_id
       FROM t_commercial_invoice`)
    defer rows.Close()
    checkErr(err)
 
    var records []*string
    for rows.Next() {
        p := new(string)
        if err := rows.Scan(&p); err != nil {
            log.Printf("sql error")
        }
        records = append(records, p)
    }
    for _,i:=range records{
        credit_start_by_company_id(*i)
    }
    fmt.Fprintln(w, "finish")
}
func credit_start_by_company_id(company_id string) {
    rows, err := db.Query(`SELECT 
       company_id,
       invoice_no,
       invoice_date,
       sales_order_id,
       payment_dead_line,
       payment_due,
       status  FROM t_commercial_invoice where company_id=?`,company_id)
    defer rows.Close()
    checkErr(err)
 
    var records []*commercial_invoice
    for rows.Next() {
        p := new(commercial_invoice)
        if err := rows.Scan(&p.company_id,
          &p.invoice_no,
          &p.invoice_date,
          &p.sales_order_id,
          &p.payment_dead_line,
          &p.payment_due,
          &p.status); err != nil {
            log.Printf("sql error")
        }
        records = append(records, p)
    }

    // print_invoice(records)
    save_to_credit_black_list(records)
    
    deal_with_sales_order()
    deal_with_statistic_time()
    print_credit_black_list(g_credit_black_list)
    //need to sort before insert to database
    sort_list()
    //delete same company_id and accouting_period_id before insert
    delete_same(g_credit_black_list[0].company_id,g_credit_black_list[0].accounting_period_id)
    insert_to_database()

    //insert to t_credit_black_list_detail
    insert_to_detail()
    //在这里清除g_credit_black_list
    fmt.Println(len(g_credit_black_list))
    g_credit_black_list = nil
    fmt.Println(len(g_credit_black_list))
}
func sort_list() {
    sort.Sort(g_credit_black_list)
    for index,i:=range g_credit_black_list{
        i.sort_no=index+1
    }
    //删除超过10个的元素
    if(len(g_credit_black_list)>10){
        g_credit_black_list=g_credit_black_list[0:10]
    }
}
func delete_same(company_id,accounting_period_id string) {
    _, err := db.Exec(
    `delete from t_credit_black_list where company_id=? and accounting_period_id=?`,company_id,accounting_period_id)
     checkErr(err)
}
func insert_to_detail() {
    for _,i:=range g_credit_black_list{
    _, err := db.Exec(
        `INSERT INTO t_credit_black_list_detail (
       detail_id,
       credit_black_list_id,
       overdue_sales_order_id,
       overdue_days,
       overdue_days_beginning_date,
       overdue_amount,
       currency_id,
       createAt,
       createBy,
       dr,
       data_version) VALUES (?, ?,?, ?,?, ?,?, ?,?, ?,?)`,rand_string(20),
        i.credit_black_list_id,i.sales_order_id,i.total_overdue_days,i.payment_dead_line,i.total_overdue_amount,i.currency_id,i.createAt,i.createBy,i.dr,i.data_version)
        checkErr(err)
    }
}
func insert_to_database() {
    for _,i:=range g_credit_black_list{
    _, err := db.Exec(
    `INSERT INTO t_credit_black_list (
        credit_black_list_id,
       company_id,
       customer_master_id,
       status,
       total_overdue_days,
       total_overdue_amount,
       currency_id,
       statistic_beginning_date,
       statistic_ending_date,
       accounting_period_id,
       sort_no,
       createAt,
       createBy,
       dr,
       data_version) VALUES (?, ?,?, ?,?, ?,?, ?,?, ?,?, ?,?, ?,?)`,
        i.credit_black_list_id,i.company_id,i.customer_master_id,i.status,i.total_overdue_days,i.total_overdue_amount,i.currency_id,i.statistic_beginning_date,i.statistic_ending_date,i.accounting_period_id,i.sort_no,i.createAt,i.createBy,i.dr,i.data_version)
        checkErr(err)
    }
}
func deal_with_statistic_time() {
    t_now:=time.Now()
     for _,i:=range g_credit_black_list{
        err := db.QueryRow("select opening_date,ending_date, accounting_period_id from t_accounting_period where company_id=? and ?>opening_date and ?<ending_date",i.company_id,t_now,t_now).Scan(&i.statistic_beginning_date,&i.statistic_ending_date,&i.accounting_period_id)
        checkErr(err)
    }
}
func save_to_credit_black_list( records []*commercial_invoice) {
    for _,i:=range records{
        p := new(credit_black_list)
        p.company_id=i.company_id
        p.sales_order_id=i.sales_order_id
        p.status=i.status
        p.dr=0
        p.data_version=1
        p.credit_black_list_id=rand_string(20)

        t_now:=time.Now()
        p.createAt=t_now.Format("2006-01-02 15:04:05")
        p.createBy="data_analysis"
        const TimeFormat = "2006-01-02"
        // t_payment_dead_line, err := time.Parse(TimeFormat, i.payment_dead_line+" CST")
        loc, _ := time.LoadLocation("Local")
        t_payment_dead_line, err:= time.ParseInLocation(TimeFormat, i.payment_dead_line, loc)
        if err != nil {
            log.Printf("time parse error", err.Error())
        }
        dur:=t_now.Sub(t_payment_dead_line)
        fmt.Println(t_now)
        // fmt.Println(i.payment_dead_line)
        fmt.Println(t_payment_dead_line)
        p.payment_dead_line=i.payment_dead_line
        fmt.Println(int(dur.Hours()/24+1))
        p.total_overdue_days=int(dur.Hours()/24+1)
        
        p.total_overdue_amount=i.payment_due
        g_credit_black_list = append(g_credit_black_list, p)

    }
}
func deal_with_sales_order() {
    for _,i:=range g_credit_black_list{
        // var customer_master_id string
        // var currency_id string
        err := db.QueryRow("select customer_master_id,currency_id from t_sales_order where sales_order_id=?",i.sales_order_id).Scan(&i.customer_master_id,&i.currency_id)
        if err != nil {
            log.Printf("sql error", err.Error())
        }
    }
}
func print_invoice( records []*commercial_invoice) {
    for _,i:=range records{
        i.print()
    }
}
func print_credit_black_list( records []*credit_black_list) {
    for _,i:=range records{
        i.print()
    }
}