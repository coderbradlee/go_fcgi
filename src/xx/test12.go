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
	"reflect"
	// "unsafe"
	"os"
	"runtime.pprof"
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
	cpu,_:=os.Create("cpu.out")
	defer cpu.Close()
	pprof.StartCPUProfile(cpu)
	defer pprof.StopCPUProfile()
	mem,_:=os.Create("mem.out")
	defer mem.Close()
	defer pprof.WriteHeapProfile(mem)
}

