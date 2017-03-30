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
type Integer int
func (a Integer)Less(b Integer)bool {
	return a<b
}
func (a *Integer)Add(b Integer) {
	*a+=b
}
type LessAdder interface{
	Less(b Integer)bool
	Add(b Integer)
}

func main() {
	var a Integer=1
	var b LessAdder=&a
	b.Add(1)
	fmt.Println(a)
	switch b.(type){
		case LessAdder:fmt.Println("int")
		default:fmt.Println("default")
	}
	
}


