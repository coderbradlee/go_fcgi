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
	// "encoding/json"
	// "bytes"
	// "os/exec"
	// "strings"
	// "syscall"
	// "log"
	// "net"
	// "net/http"
	"errors"
	"net/rpc"
)

var lock sync.Mutex

func init() {
	// runtime.NumCPU()
	runtime.GOMAXPROCS(runtime.NumCPU())
}

type Args struct{
	A,B int
}
type Quotient struct{
	Quo,Rem int
}
type Arith int
func (t *Arith)Multiply(args *Args,reply *int)error {
	*reply=args.A*args.B
	return nil
}
func (t *Arith)Divide(args *Args,quo *Quotient)error {
	if args.B==0{
		return errors.New("divide by zero")
	}
	quo.Quo=args.A/args.B
	quo.Rem=args.A%args.B
	return nil
}
func main() {
	client,err:=rpc.DialHTTP("tcp","127.0.0.1:1234")
	if err!=nil{
		fmt.Println("dialing:",err)
	}
	args:=&Args{17,8}
	var reply int
	err=client.Call("Arith.Multiply",args,&reply)
	if err!=nil{
		fmt.Println("arith error:",err)
	}
	fmt.Printf("%d*%d=%d\n",args.A,args.B,reply)
	quotient:=new(Quotient)
	divCall:=client.Go("Arith.Divide",args,&quotient,nil)
	<-divCall.Done
	fmt.Println(err)
	fmt.Println(quotient.Quo)
	fmt.Println(quotient.Rem)
	// fmt.Println(r.Reply)
}
