package main
//#go:generate ls -l
import (
	"fmt"
	// "regexp"
	// "os"
	// "bufio"
	"runtime"
	"io"
	"io/ioutil"
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
)

var lock sync.Mutex

func init() {
	// runtime.NumCPU()
	runtime.GOMAXPROCS(runtime.NumCPU())
}
func bind(f func(in io.Reader,out io.Writer,p []string),params []string)func(in io.Reader,out io.Writer) {
	return func(in io.Reader,out io.Writer) {
		f(in,out,params)
	}
}
func pipe(app1 func(in io.Reader,out io.Writer),app2 func(in io.Reader,out io.Writer))func(in io.Reader,out io.Writer) {
	return func(in io.Reader,out io.Writer) {
		r,w:=os.Pipe()
		defer w.Close()
		app1(in,w)
		go func() {
			defer r.Close()
			app2(r,out)
		}
	}
}
func main() {
	f:=func(in io.Reader,out io.Writer,params []string) {
		var command string
		if b, err := ioutil.ReadAll(in); err == nil {
		    command=string(b)
		}
		cmd := exec.Command(command, params)
		cmd.Stdout =out
		err = cmd.Start()
		if err != nil {
			fmt.Println("cmd error!")
		}
	}
	fp:=bind(f,"note")
	// fp2:=bind(f,"select")
	fp(strings.NewReader("cat"),os.Stdout)

	// pp:=pipe(fp,fp2)
	// pp(strings.NewReader("cat"),os.Stdout)//cat note|grep select
	                      //
	fmt.Println("done!")
}


