package main

import (
	"io"
	"log"
	"path"
	"net/http"
	"io/ioutil"
	"html/template"
	"runtime/debug"
)
const(
	ListDir=0x0001
	UPLOAD_DIR="./uploads"
	TEMPLATE="./views"
	templates:=make(map[string]*template.Template)

)

var lock sync.Mutex

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fileInfoAdd,err:=ioutil.ReadDir(TEMPLATE_DIR)
	check(err)
	var templateName,templatePath string
	for _,fileInfo:=range fileInfoAdd{
		templateName=fileInfo.Name
		
	}
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
