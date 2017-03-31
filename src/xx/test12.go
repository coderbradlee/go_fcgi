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
	"time"
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
type Vector []float64
func (v *Vector)doSome(i,n int,u Vector,c chan<- int) {
	for ;i<n;i++{
		v[i]+=u[i]
	}
	c<-1
}
func (v *Vector)doAll(u Vector) {
	c:=make(chan int,4)
	for i:=0;i<4;i++{
		go v.doSome(i*len(v)/4,(i+1)*len(v)/4+1,u,c)
	}
	for i:=range c{
		fmt.Println(i)
	}
}
func main() {
	var v =Vector{1,2,3,4,5,6,7,8,9,10}
	var u =Vector{10,9,8,7,6,5,4,3,2,1}
	v.doAll()
}


