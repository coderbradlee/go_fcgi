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
	done:=make(chan struct{})
	for i:=0;i<3;i++{
		go func (index int) {
			select{
				case ch<-(index+1)*2:fmt.Println("index:",index)
				case <-done:fmt.Println("done:",index)
			}	
		}(i)
	}
	fmt.Println("1:",<-ch)
	// fmt.Println("2:",<-ch)
	// fmt.Println("3:",<-ch)
	close(done)
	time.Sleep(2*time.Second)
}

