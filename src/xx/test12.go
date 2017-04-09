package main

//#go:generate ls -l
import (
	"fmt"
	// "regexp"
	// "os"
	// "bufio"
	"io"
	//	"io/ioutil"
	"runtime"
	"sync"
	// "testing"
	// "math"
	// "reflect"
	// "unsafe"
	//	"os"
	// "runtime/pprof"
		"time"
	// "encoding/json"
	"bytes"
	// "os/exec"
	// "strings"
	// "syscall"
	// "log"
	"net"
	"net/http"
	"reflect"
)

var lock sync.Mutex

func init() {
	// runtime.NumCPU()
	runtime.GOMAXPROCS(runtime.NumCPU())
}

var once sync.Once
var a string

func setup() {
	a = "x.x.x.x:8088"
}
func do() {
	once.Do(setup)
	fmt.Println(a)
}
func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
func checkSum(msg []byte) uint16 {
	sum := 0
	for n := 1; n < len(msg)-1; n += 2 {
		sum += int(msg[n])*256 + int(msg[n+1])
	}
	sum = (sum >> 16) + (sum & 0xffff)
	sum += (sum >> 16)
	var answer uint16 = uint16(^sum)
	return answer
}
func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()
	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				fmt.Println("break:", n)
				break
			}
			return nil, err
		}
	}
	return result.Bytes(), nil
}
type customTransport struct{
	Transport http.RoundTripper
}
func (t *customTransport)transport()http.RoundTripper {
	if t.Transport!=nil{
		return t.Transport
	}
	return http.DefaultTransport
}
func (t *customTransport)RoundTrip(req *http.Request)(*http.Response,error) {
	fmt.Println("RoundTrip")
	return t.transport().RoundTrip(req)
}
func (t *customTransport)Client()*http.Client {
	return &http.Client{Transport:t}
}
func testreflect() {
	var x float64=3.4
	// v:=reflect.ValueOf(x)
	// fmt.Println("type:",v.Type())
	// fmt.Println("kind:",v.Kind())
	// fmt.Println("value:",v.Float())
	// var y float64=4.1
	// v.Set(y)
	// fmt.Println(v.CanSet())
	p:=reflect.ValueOf(&x)
	fmt.Println(p.Type())
	fmt.Println(p.CanSet())
	v:=p.Elem()
	fmt.Println(v.CanSet())
	v.SetFloat(7.1)
	fmt.Println(v.Interface())
	fmt.Println(x)
}
func testreflect1() {
	type T struct{
		A int
		B string
	}
	t:=T{203,"mh203"}
	s:=reflect.ValueOf(&t).Elem()
	typeoft:=s.Type()
	for i:=0;i<s.NumField();i++{
		f:=s.Field(i)
		fmt.Printf("%d:%s %s=%v\n",i,typeoft.Field(i).Name,f.Type(),f.Interface())
	}

}
type ball struct{
	hit int
}
func player(name string,b chan *ball) {
	for{
		xx:=<-b
		fmt.Println(name,xx.hit)
		xx.hit++
		time.Sleep(100*time.Millisecond)
		b<-xx
	}	
}
func test_pingpong() {
	table:=make(chan *ball)
	table<-new(ball)
	fmt.Println("toss a ball")
	<-table
	fmt.Println("throw a ball")
	go player("ping",table)
	go player("pong",table)
	table<-new(ball)
	time.Sleep(1*time.Second)
	<-table
}
func main() {
	test_pingpong()
	// testreflect()
	// testreflect1()
	//	service := "www.qq.com:80"
	//	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	//	checkError(err)
	//	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	//	checkError(err)
	//	fmt.Println(tcpAddr)
	//	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	//	checkError(err)
	//	result, err := ioutil.ReadAll(conn)
	//	checkError(err)
	//	fmt.Println(string(result))
	//	///////////////////////////////
	//	domain := "www.qq.com"
	//	addrs, err := net.LookupHost(domain)
	//	//	fmt.Println(cname)
	//	fmt.Println(addrs)
	//	fmt.Println(err)
	// resp, _ := http.Head("https://www.baidu.com/")
	// fmt.Println(resp)
	// t:=&customTransport{}
	// c:=t.Client()
	// _,err:=c.Get("https://www.baidu.com/")
	// fmt.Println(err)
	// b := []byte(`{
	// "Title": "Go语言编程",
	// "Authors": ["XuShiwei", "HughLv", "Pandaman", "GuaguaSong", "HanTuo", "BertYuan",
	// "XuDaoli"],
	// "Publisher": "ituring.com.cn",
	// "IsPublished": true,
	// "Price": 9.99,
	// "Sales": 1000000
	// }`)
	// var r interface{}
	// err := json.Unmarshal(b, &r)
	// if err!=nil{
	// 	fmt.Println(err)
	// }
	// if gobook,ok:=r.(map[string]interface{});ok{
	// 	fmt.Println("map[string]interface{}")
	// 	for k,v:=range gobook{
	// 		fmt.Println(k,":",v)
	// 	}
	// }
}
