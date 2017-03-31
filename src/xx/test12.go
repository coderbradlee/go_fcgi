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
	// "os"
	// "runtime/pprof"
	// "time"
	// "encoding/json"
	// "bytes"
	// "os/exec"
	// "strings"
	// "syscall"
	// "log"
)

var lock sync.Mutex

func init() {
	// runtime.NumCPU()
	runtime.GOMAXPROCS(runtime.NumCPU())
}
var  counter int=0
func Count(ch chan int) {

	counter++
	fmt.Println(counter)
	ch<-counter
}
func Parse(ch <-chan int) {
	for v:=range ch{
		fmt.Println("parse:",v)
	}
	// <-ch
}
func main() {
	ch:=make(chan int,2)
	//for{
		select{
		case ch<-0:fmt.Println("0")
		case ch<-1:fmt.Println("1")
		}
		go Parse(ch)
		// <-ch
		close(ch)

	//}
	time.Sleep(2*time.Second)
}


