package main
 
import (
    "fmt"
    "log"
    "net/http"
    "os"
    "encoding/json"
    _"time"
    _"strings"
    _"io/ioutil"
    "flag"
    "runtime" 
    "net" 
    "net/http/fcgi"
    "martini"
    "database/sql")
type mysql_conf struct{
    Host string
    Port string
    Username string
    Password string
    Database string
}
type Configuration struct {
    Exec_time string
    FastcgiPort string
    Log_name string
    HttpPort string
    RedisNodes []string
    Mysql_conf mysql_conf
}
var configuration Configuration
var db *sql.DB
func init() {
    runtime.GOMAXPROCS(runtime.NumCPU())
    file, _ := os.Open("src/mainproject/conf.json")
    decoder := json.NewDecoder(file)
    configuration = Configuration{}
    err := decoder.Decode(&configuration)
    if err != nil {
      fmt.Println("error:", err)
    }
    fmt.Printf("Exec_time %s\n",configuration.Exec_time)
    fmt.Printf("FastcgiPort %s\n",configuration.FastcgiPort)
    fmt.Printf("Log_name %s\n",configuration.Log_name)
    fmt.Printf("HttpPort %s\n",configuration.HttpPort)
    mysql_init()
    log_init()
    //fmt.Println(configuration.exec_time) // output: [UserA, UserB]
}
func mysql_init() {
    var conn_string string
    conn_string=configuration.Mysql_conf.Username+":"+configuration.Mysql_conf.Password+"@tcp("+configuration.Mysql_conf.Host+":"+configuration.Mysql_conf.Port+")/"+configuration.Mysql_conf.Database+"?charset=utf8"
    // db, _ = sql.Open("mysql", "renesola:renes0la.xx@tcp(172.18.22.202:3306)/apollo_eu_erp?charset=utf8")
    fmt.Printf("conn_string:%s\n",conn_string)
    var err error
    db, err = sql.Open("mysql", conn_string)
    if err != nil {
        panic("dbpool init >> " + err.Error())
    }
    db.SetMaxOpenConns(20)
    db.SetMaxIdleConns(10)
    db.Ping()
}
func log_init() {
    log_name:=fmt.Sprintf("%s",configuration.Log_name)
    logFileName := flag.String("log", log_name, "Log file name")
    
    flag.Parse()

    //set logfile Stdout
    logFile, logErr := os.OpenFile(*logFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
    if logErr != nil {
        fmt.Println("Fail to find", *logFile, "cServer start Failed")
        os.Exit(1)
    }
    log.SetOutput(logFile)
    log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}


func main() {
    // go startHttpServer()
    go startMartini()
    port:=fmt.Sprintf("%s",configuration.FastcgiPort)
    l, err := net.Listen("tcp", port)
    if err != nil { 
        panic(err) 
    } 
    serveMux := http.NewServeMux() 
    serveMux.HandleFunc("/redis", redisHandler) 
    serveMux.HandleFunc("/pdf", pdfHandler) 
    err = fcgi.Serve(l, serveMux)
    if err != nil { 
        log.Println("fcgi error", err) 
    }
    // go startHttpServer()
}
func startHttpServer() {
    port:=fmt.Sprintf("%s",configuration.HttpPort)
    http.HandleFunc("/redis", redisHandler) 
    http.HandleFunc("/pdf", pdfHandler) 
    // http.HandleFunc("/po/deliver_goods",poHandler)
    err := http.ListenAndServe(port, nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
func startMartini() {
    port:=fmt.Sprintf("%s",configuration.HttpPort)
    m := martini.Classic()
    m.Post("/po/deliver_goods",poHandler)
    m.RunOnAddr(port)
    ////////////////////

    // log_name:=fmt.Sprintf("%s",configuration.Log_name)
    // logFileName := flag.String("log", log_name, "Log file name")
    
    // flag.Parse()

    // //set logfile Stdout
    // logFile, logErr := os.OpenFile(*logFileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
    // if logErr != nil {
    //     fmt.Println("Fail to find", *logFile, "cServer start Failed")
    //     os.Exit(1)
    // }
    // log.SetOutput(logFile)
    // log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

///////////////////////////////
    // l,err:=GetLogger("log")
    // if err != nil {
    //     fmt.Println("Fail to find logFile cServer start Failed")
    //     os.Exit(1)
    // }
    // m.Logger(l)
    m.Run()
}
