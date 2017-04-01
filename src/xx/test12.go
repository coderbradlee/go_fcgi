package main
//#go:generate ls -l
import (
	"fmt"
	// "regexp"
	// "os"
	// "bufio"
	"runtime"
	// "io"
	// "io/ioutil"
	"sync"
	// "testing"
	// "math"
	// "reflect"
	// "unsafe"
	"os"
	// "runtime/pprof"
	// "time"
	// "encoding/json"
	// "bytes"
	// "os/exec"
	// "strings"
	// "syscall"
	// "log"
	"net"

)

var lock sync.Mutex

func init() {
	// runtime.NumCPU()
	runtime.GOMAXPROCS(runtime.NumCPU())
}
var once sync.Once
var a string
func setup() {
	a="x.x.x.x:8088"
}
func do() {
	once.Do(setup)
	fmt.Println(a)
}
func checkError(err error) {
	if err!=nil{
		fmt.Println(err)
	}
}
func checkSum(msg []byte)uint16 {
	sum:=0
	for n:=1;n<len(msg)-1;n+=2{
		sum+=int(msg[n])*256+int(msg[n+1])
	}
	sum=(sum>>16)+(sum&0xffff)
	sum+=(sum>>16)
	var answer uint16=uint16(^sum)
	return answer
}
func readFully(conn net.Conn)([]byte,error) {
	defer conn.Close()
	result:=bytes.NewBuffer(nil)
	var buf [512]byte
	for{
		n,err:=conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err!=nil{
			if err==io.EOF{
				break
			}
			return nil,err
		}
	}
	return result.Bytes(),nil
}
func main() {
	service:="www.baidu.com"
	conn,err:=net.Dial("tcp",service)
	checkError(err)
	_,err=conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)
	result,err:=readFully(conn)
	checkError(err)
	fmt.Println(string(result))
	os.Exit(0)
}


