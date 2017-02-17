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
    "net/http/fcgi")
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
    fmt.Printf("%s\n",configuration.Log_name)
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
    startHttpServer()
}
func startHttpServer() {
    port:=fmt.Sprintf("%s",configuration.HttpPort)
    http.HandleFunc("/redis", redisHandler) 
    http.HandleFunc("/pdf", pdfHandler) 
    err := http.ListenAndServe(port, nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

