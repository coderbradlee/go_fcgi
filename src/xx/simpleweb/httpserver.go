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
func check(err error) {
	if err!=nil{
		fmt.Println(err)
	}
}
func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fileInfoArr,err:=ioutil.ReadDir(TEMPLATE_DIR)
	check(err)
	var templateName,templatePath string
	for _,fileInfo:=range fileInfoArr{
		templateName=fileInfo.Name
		if ext:=path.Ext(templateName;ext!=".html"){
			continue
		}
		templatePath=TEMPLATE_DIR+"/"+templateName
		fmt.Println("loading:",templatePath)
		t:=template.Must(template.ParseFiles(templatePath))
		templates[tmpl]=t
	}
}
func renderHtml(w http.ResponseWriter,tmpl string,locals map[string]interface{}) {
	err:=templates[tmpl].Execute(w,locals)
	check(err)
}
func isExists(path string)bool {
	_,err:=os.Stat(path)
	if err==nil{
		return true
	}
	return os.IsExist(err)
}
func uploadHandler (w http.ResponseWriter, r *http.Request) {
	if r.Method=="GET"{
		renderHtml(w,"upload",nil)
	}
	if r.Method=="POST"{
		f,h,err:=r.FormFile("image")
		check(err)
		filename:=h.Filename
		defer f.Close()
		t,err:=ioutil.TempFile(UPLOAD_DIR,filename)
		check(err)
		defer t.Close()
		_,err:=io.Copy(t,f)
		check(err)
		http.Redirect(w,r,"/view?id="+filename,http.StatusFound)
	}

}
func viewHandler(w http.ResponseWriter,r *http.Request) {
	imageId:=r.FormValue("id")
	imagePath:=UPLOAD_DIR+"/"+imageId
	if exists:=isExists(imagePath);!exists{
		http.NotFound(w,r)
		return
	}
	w.Header.Set("Content-Type","image")
	http.ServeFile(w,r,imagePath)
}
func listHandler(w http.ResponseWriter,r *http.Request) {
	fileInfoArr,err:=ioutil.ReadDir("./uploads")
	check(err)
	locals:=make(map[string]interface{})
	images:=[]string{}
	for _,fileInfo:=range fileInfoArr{
		images=append(images,fileInfo.Name)
	}
	locals["images"]=images
	renderHtml(w,"list",locals)
}
func safeHandler(fn http.HandlerFunc)http.HandlerFunc {
	return func(w http.ResponseWriter,r *http.Request){
		defer func() {
			if e,ok:=recover().(error);ok{
				http.Error(w,e.Error(),http.StatusInternalServerError)
				fmt.Printf("panic:%v - %v\n",fn,e)
				fmt.Println(string(debug.Stack()))
			}
		}()
		fn(w,r)
	}
}
func staticDirHandler(mux *http.ServeMux,prefix string,staticDir string,flags int) {
	mux.HandlerFunc(prefix,func(w http.ResponseWriter,r *http.Request) {
		file:=staticDir+r.URL.Path[len(prefix)-1:]
		if(flags&ListDir)==0{
			if exists:=isExists(file);!exists{
				http.NotFound(w,r)
				return
			}
		}
		http.ServeFile(w,r,file)
	})
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
	go func() {
		mux := http.NewServeMux()
		staticDirHandler(mux, "/assets/", "./public", 0)
		mux.HandleFunc("/", safeHandler(listHandler))
		mux.HandleFunc("/view", safeHandler(viewHandler))
		mux.HandleFunc("/upload", safeHandler(uploadHandler))
		err := http.ListenAndServe(":8080", mux)
		if err != nil {
			fmt.Println("ListenAndServe: ", err.Error())
		}
	}
    http.HandleFunc("/json", jsonHandler)   
    http.ListenAndServe(":8877", nil)
}
