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
	// "encoding/json"
	// "bytes"
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
func get() []byte {  
    raw := make([]byte,10000)
    fmt.Println(len(raw),cap(raw),&raw[0]) //prints: 10000 10000 <byte_addr_x>
    return raw[:3]
}
type myLocker struct{
	sync.Mutex
}
type field struct{
	name string
}
type data struct {  
    name string
}
func main() {
	s := []data
    s[0].name = "two" //ok
    fmt.Println(s)    //prints: [{two}]
}


