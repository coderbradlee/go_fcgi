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
    "martini")
type Configuration struct {
    Exec_time string
    FastcgiPort string
    Log_name string
    HttpPort string
    RedisNodes []string
}
var configuration Configuration

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
    log_init()
    //fmt.Println(configuration.exec_time) // output: [UserA, UserB]
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
    // l:=log.Logger
    m.Logger(log)
    m.Run()
}
