package main
import (
	"fmt"
	"net/http"
)
const(
	SERVER_PORT=8080
	SERVER_DOMAIN="xx.com"
	RESPONSE_TEMPLATE="HELLO"
)
func rootHandler(w http.ResponseWriter,r *http.Request) {
	w.Header().Set("Content-Type","text/html")
	w.Header().Set("Content-Length",fmt.Sprint(len(RESPONSE_TEMPLATE)))
	w.Write([]byte(RESPONSE_TEMPLATE))
}
func main() {
	// http.HandleFunc(fmt.Sprintf("%s:%d/",SERVER_DOMAIN,SERVER_PORT),rootHandler)
	h := http.FileServer(http.Dir("."))
	err:=http.ListenAndServeTLS(fmt.Sprintf(":%d",SERVER_PORT),"/root/cert/server.crt","/root/cert/server.key",h)
	if err!=nil{
		fmt.Println(err)
	}
}
