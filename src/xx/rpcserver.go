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
	//	"time"
	// "encoding/json"
	"bytes"
	// "os/exec"
	// "strings"
	// "syscall"
	// "log"
	"net"
	"net/http"
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

func main() {
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
	t:=&customTransport{}
	c:=t.Client()
	_,err:=c.Get("https://www.baidu.com/")
	fmt.Println(err)
}
