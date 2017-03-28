package main
//#go:generate ls -l
import (
	"fmt"
	// "regexp"
	// "os"
	// "bufio"
	"runtime"
	// "io"
	"sync"
	// "testing"
	// "math"
	// "reflect"
	// "unsafe"
	// "os"
	// "runtime/pprof"
	"time"
)
type Data struct{
	num int
	key *string
	items map[string]bool
}
func (this *Data)pmethod() {
	this.num=7
}
func (this Data)vmethod() {
	this.num=8
	*this.key="v.key"
	this.items["vmethod"]=true
}
var lock sync.Mutex
func test(){
	lock.Lock()
	lock.Unlock()
}
func test_defer() {
	lock.Lock()
	defer lock.Unlock()
}

func init() {
	// runtime.NumCPU()
	runtime.GOMAXPROCS(runtime.NumCPU())
}
func main() {
	key:="key1"
	d:=Data{1,&key,make(map[string]bool)}
	fmt.Printf("num=%v key=%v items=%v\n",d.num,*d.key,d.items)
	d.pmethod()
	fmt.Printf("num=%v key=%v items=%v\n",d.num,*d.key,d.items)
	d.vmethod()
	fmt.Printf("num=%v key=%v items=%v\n",d.num,*d.key,d.items)
}

