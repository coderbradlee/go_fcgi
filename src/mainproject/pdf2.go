package main

import (
    // _"log"
    "fmt"
    // // "net/http"
    // "unsafe"
    "encoding/json"
    "net/http"
    "io/ioutil"
    "logger"
    // "strconv"
    // "errors"
    // "os"
    // "converter"
    // "reflect"
    "os/exec"
)
type src_dst struct{
	Src string
	Dst string
}
func pdfHandler2 (w http.ResponseWriter, r *http.Request) {

  	addr := r.Header.Get("X-Real-IP")
	if addr == "" {
		addr = r.Header.Get("X-Forwarded-For")
		if addr == "" {
			addr = r.RemoteAddr
		}
	}
  
	body, _:= ioutil.ReadAll(r.Body)
	// logger.Info(fmt.Sprintf("%s",body))
	defer r.Body.Close()
    var ret string
    var t src_dst  
    err_decode := json.Unmarshal(body, &t)
    if err_decode!=nil{
    	ret=`decode failed`
	    fmt.Fprint(w,ret )
	    logger.Error(fmt.Sprintf("Started %s %s for %s:%s response:%s", r.Method, r.URL.Path, addr,body,err_decode.Error()))
	    return
    }
    // logger.Info(t.Src+": "+t.Dst)
    if t.Src==""||t.Dst==""{
    	ret="empty src or dst"
    	fmt.Fprint(w,ret )
    	logger.Error(fmt.Sprintf("Started %s %s for %s:%s response:%s", r.Method, r.URL.Path, addr,body,"empty src or dst"))
	    return
    }
    var err error
    err=convert2(t.Src,t.Dst)
    if err!=nil{
		fmt.Fprint(w,err.Error())
		
    	logger.Error(fmt.Sprintf("Started %s %s for %s:%s response:%s", r.Method, r.URL.Path, addr,body,err.Error()))
    	return
    }else{
    	fmt.Fprint(w,"ok")
    }
    logger.Info(fmt.Sprintf("Started %s %s for %s:%s response:%s", r.Method, r.URL.Path, addr,body,"ok"))
} 
func convert2(src,dst string) error{
	logger.Info("convert2")
	var err error
	cmd := exec.Command("wkhtmltopdf", src,dst)
	err = cmd.Start()
	if err != nil {
		logger.Error(err)
	}
	return err
}
