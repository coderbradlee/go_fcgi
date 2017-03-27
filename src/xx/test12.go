package main
//#go:generate ls -l
import (
	"fmt"
	"regexp"
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
)
type Data struct{

}
func (*Data)String() string{
	return ""
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
	payment_type:="30 Days xxxxxx"
	reg := regexp.MustCompile(`([1-9][0-9])( Days)`)
    reg_find:=reg.FindAllString(payment_type, -1)[0]
    fmt.Printf("%s: %s\n",payment_type, reg_find[:2])

}

