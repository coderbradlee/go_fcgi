package main

//#go:generate ls -l
import (
	"fmt"
	// "regexp"
	// "os"
	// "bufio"
	// "io"
	//	"io/ioutil"
	"runtime"
	"sync"
	// "testing"
	// "math"
	// "reflect"
	// "unsafe"
	//	"os"
	// "runtime/pprof"
	//	"time"
	"encoding/json"
	// "bytes"
	// "os/exec"
	// "strings"
	// "syscall"
	// "log"
	// "net"
	"net/http"
	// "errors"
	// "net/rpc"
	// "martini"
)

var lock sync.Mutex

func init() {
	// runtime.NumCPU()
	runtime.GOMAXPROCS(runtime.NumCPU())
}
type PoData struct{
	Request_system int32 `json:"request_system"`
	Request_time string `json:"request_time"`
   	Operation string `json:"operation"`
}
func jsonHandler (w http.ResponseWriter, r *http.Request) {
	////////////////////////////////
	addr := r.Header.Get("X-Real-IP")
	if addr == "" {
		addr = r.Header.Get("X-Forwarded-For")
		if addr == "" {
			addr = r.RemoteAddr
		}
	}
/////////////////////////////////////////////////////////////////
	if r.Method !="POST"{
		fmt.Fprint(w, "this interface should be post!")
	} else{
		decoder := json.NewDecoder(r.Body)
	    var t PoData
		defer r.Body.Close()
	    err_decode := decoder.Decode(&t)

	    enc := json.NewEncoder(w)

	    if err_decode != nil {
	        fmt.Println(err_decode)
	        return;
	    }
	    if err := enc.Encode(&t); err != nil {
			fmt.Println(err)
		}
	    // fmt.Fprint(w,ret )
	}

} 
func main() {
    http.HandleFunc("/json", jsonHandler)   
    http.ListenAndServe(":8877", nil)
}
