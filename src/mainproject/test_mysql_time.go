 package main
 import (
    "logger"
    "fmt"
    // "encoding/json"
    "net/http"
    // "io/ioutil"
    // "bytes"
    // "time"
    // "errors"
    // "runtime/pprof"
)
func test_mysql_time (w http.ResponseWriter, r *http.Request) {
	////////////////////////////////
	addr := r.Header.Get("X-Real-IP")
	if addr == "" {
		addr = r.Header.Get("X-Forwarded-For")
		if addr == "" {
			addr = r.RemoteAddr
		}
	}
	if r.Method !="POST"{
		fmt.Fprint(w, "this interface should be post!")
	} else{
		ret:=single_select()
	    fmt.Fprint(w,ret )
	    log_str:=fmt.Sprintf("Started %s %s for %s:%s response:%s", r.Method, r.URL.Path, addr,"body",ret)
        logger.Info(log_str)
	}

} 

func single_select()string {
	var packing_method string
    db.QueryRow("select packing_method_id from t_packing_method where name=?","Pallet").Scan(&packing_method)
    return packing_method
}