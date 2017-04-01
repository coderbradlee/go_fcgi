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
func main() {
	service:="127.0.0.1"
	conn,err:=net.Dial("ip4:icmp",service)
	checkError(err)
	var msg [512]byte
	msg[0]=8
	msg[1]=0
	msg[2]=0
	msg[3]=0
	msg[4]=0
	msg[5]=13
	msg[6]=0
	msg[7]=37
	len:=8
	check:=checkSum(msg[0:len])
	msg[2]=byte(check>>8)
	msg[3]=byte(check&255)
	_,err=conn.Write(msg[0:len])
	checkError(err)
	_,err=conn.Read(msg[0:])
	checkError(err)
	fmt.Println("got response")
	if msg[5]==13{
		fmt.Println("13 ok")
	}
	if msg[7]==37{
		fmt.Println("37 ok")
	}
	for i:=0;i<8;i++{
		fmt.Printf("%d %d\n",i,msg[i])
	}
	os.Exit(0)
}


