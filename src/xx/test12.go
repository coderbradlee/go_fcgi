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
	ch:=make(chan int)
	for i:=0;i<3;i++{
		go func (index int) {
			ch<-(index+1)*2
		}(i)
	}
	fmt.Println("1:",<-ch)
	fmt.Println("2:",<-ch)
	fmt.Println("3:",<-ch)
	close(ch)
	time.Sleep(2*time.Second)
}

