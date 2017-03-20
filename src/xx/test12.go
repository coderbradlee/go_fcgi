package main
#go:generate ls -l
import (
	"fmt"
	// "regexp"
	// "os"
	// "bufio"
	"runtime"
	// "io"
	// "sync"
	// "testing"
	// "math"
	"reflect"
	// "unsafe"
)
type Data struct{

}
func (*Data)String() string{
	return ""
}

func init() {
	// runtime.NumCPU()
	runtime.GOMAXPROCS(runtime.NumCPU())
}
func main() {
	var d *Data
	t:=reflect.TypeOf(d)
	fmt.Println(t)
	it:=reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	fmt.Println(it)
	fmt.Println(t.Implements(it))
}

