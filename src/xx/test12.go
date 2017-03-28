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
	// "time"
	// "encoding/json"
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
	defer func () {
		fmt.Println("recovered:",recover())
	}()
	panic("fail")
}

