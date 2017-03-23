package main
 
import (
    "fmt"
    "logger"
    "net/http"
    "os"
    "encoding/json"
    _"time"
    _"strings"
    _"io/ioutil"
    // "flag"
    "runtime" 
    "net" 
    "net/http/fcgi"
    "martini"
    "database/sql"
    _"mysql"
    "strconv"
    "time"
    "net/http/pprof"
    )
type mysql_conf struct{
    Host string
    Port string
    Username string
    Password string
    Database string
}
type log_conf struct{
    Dir string
    Name string
    Console bool
    Num int32
    Size int64
    Level string
}
type Configuration struct {
    Exec_time string
    FastcgiPort string
    Log log_conf
    HttpPort string
    RedisNodes []string
    Mysql_conf mysql_conf
    Redis_url string
    System_no string
    Nfs_path string
}
var configuration Configuration
var db *sql.DB
var cpuprofile ="go_fcgi.prof"
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
    // fmt.Printf("Log_name %s\n",configuration.Log_name)
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
    // log_name:=fmt.Sprintf("%s",configuration.Log.Name)
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
    // logger.SetConsole(true)
    logger.SetConsole(configuration.Log.Console)
    logger.SetRollingFile(configuration.Log.Dir, configuration.Log.Name, configuration.Log.Num, configuration.Log.Size, logger.KB)
    //ALL，DEBUG，INFO，WARN，ERROR，FATAL，OFF
    logger.SetLevel(logger.ERROR)
    if configuration.Log.Level=="info"{
        logger.SetLevel(logger.INFO)
        }else if configuration.Log.Level=="error"{
            logger.SetLevel(logger.ERROR)
        }
}
func logPanics(function func(http.ResponseWriter,
    *http.Request)) func(http.ResponseWriter, *http.Request) {
    return func(writer http.ResponseWriter, request *http.Request) {
        defer func() {
            if x := recover(); x != nil {
                logger.Error(fmt.Sprintf("[%v] caught panic: %v", request.RemoteAddr, x))
            }
        }()
        function(writer, request)
    }
}

func main() {
    
    // go startHttpServer()
    go startfcgi()
    // go test_log()
    time.Sleep(2000 * time.Millisecond)
    // go startHttpServer()
    // go benchmark()

    startMartini()
}
func startfcgi() {
    port:=fmt.Sprintf("%s",configuration.FastcgiPort)
    l, err := net.Listen("tcp", port)
    if err != nil { 
        panic(err) 
    } 
    serveMux := http.NewServeMux() 
    serveMux.HandleFunc("/redis", redisHandler) 
    // serveMux.HandleFunc("/pdf", pdfHandler) 
    err = fcgi.Serve(l, serveMux)
    if err != nil { 
        logger.Info("fcgi error:"+err.Error()) 
    }
}
func startHttpServer() {
    port:=fmt.Sprintf("%s",configuration.HttpPort)
    http.HandleFunc("/redis", redisHandler) 
    // http.HandleFunc("/pdf", pdfHandler) 
    // http.HandleFunc("/po/deliver_goods",poHandler)
    err := http.ListenAndServe(port, nil)
    if err != nil {
        logger.Fatal("ListenAndServe: ", err)
    }
}

func test_log() {
    
    for i := 1; i > 0; i-- {
        go func() {
            logger.Debug("Debug>>>>>>>>>>>>>>>>>>>>>>" + strconv.Itoa(i))
            logger.Info("Info>>>>>>>>>>>>>>>>>>>>>>>>>" + strconv.Itoa(i))
            logger.Warn("Warn>>>>>>>>>>>>>>>>>>>>>>>>>" + strconv.Itoa(i))
            logger.Error("Error>>>>>>>>>>>>>>>>>>>>>>>>>" + strconv.Itoa(i))
            logger.Fatal("Fatal>>>>>>>>>>>>>>>>>>>>>>>>>" + strconv.Itoa(i))
        }()
        time.Sleep(100 * time.Millisecond)
    }
    time.Sleep(15 * time.Second)
    
}
func startMartini() {
    port:=fmt.Sprintf("%s",configuration.HttpPort)
    m := martini.Classic()
    // m.Use(auth_BasicFunc(func(username, password string) bool {
    // return auth_SecureCompare(username, "admin") && auth_SecureCompare(password, "admin")
    
    // }))
    // m.Use(auth_BasicFunc(func(username, password string) bool {
    //     return username == "admin" && password == "admin"
    // }))
    // m.Post("/po/deliver_goods",logPanics(poHandler))
    m.Post("/po/submit",logPanics(po_submit_Handler))
    m.Post("/po/deliver_goods",logPanics(deliver_goods_Handler))
    //m.Post("/pdf",pdfHandler)
    // m.Post("/pdf",logPanics(pdfHandler))
    m.Post("/pdf",logPanics(pdfHandler2))
    m.Post("/test_mysql_time",logPanics(test_mysql_time))
    // m.Get("/debug/pprof", pprof.Index)
    // m.Get("/debug/pprof/cmdline", pprof.Cmdline)
    m.Get("/debug/pprof/profile", pprof.Profile)
    // m.Get("/debug/pprof/symbol", pprof.Symbol)
    // m.Post("/debug/pprof/symbol", pprof.Symbol)
    // m.Get("/debug/pprof/block", pprof.Handler("block").ServeHTTP)
    // m.Get("/debug/pprof/heap", pprof.Handler("heap").ServeHTTP)
    // m.Get("/debug/pprof/goroutine", pprof.Handler("goroutine").ServeHTTP)
    // m.Get("/debug/pprof/threadcreate", pprof.Handler("threadcreate").ServeHTTP)



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
