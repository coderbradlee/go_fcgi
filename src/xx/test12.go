package main
//#go:generate ls -l
import (
	"fmt"
	// "regexp"
	"os"
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
	"time"
	// "encoding/json"
	// "bytes"
	"os/exec"
	"strings"
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
		wg := sync.WaitGroup{}
		wg.Add(2)
		r,w:=io.Pipe()
		// exit:=make(chan int)
		defer w.Close()
		go func() {
			defer r.Close()
			defer wg.Done()
			app2(r,out)
			// close(ch)
			// exit<-1
		}()
		go func(){
			defer wg.Done()
			app1(in,w)
		}()	// <-exit
		wg.Wait()
	}
}
func main() {
	func1:=func(in io.Reader,out io.Writer,params []string) {
		var command string
		if b, err := ioutil.ReadAll(in); err == nil {
		    command=string(b)
		}
		cmd := exec.Command(command, params[0])
		cmd.Stdout=out
		err := cmd.Start()
		if err != nil {
			fmt.Println("cmd error!")
		}
	}
	params1:=[]string{"test"}
	bind1:=bind(func1,params1)

	func2:=func(in io.Reader,out io.Writer,params []string) {
		// var content string
		// if b, err := ioutil.ReadAll(in); err == nil {
		//     content=string(b)
		// }else{
		// 	fmt.Println("content:",err.Error())
		// }
		// fmt.Println("content:",content)
		if _, err := io.Copy(out, in); err != nil {
			fmt.Println("copy error:",err.Error())
		}
		// cmd := exec.Command("grep", params[0],content)
		// cmd.Stdout =out
		// err := cmd.Start()
		// if err != nil {
		// 	fmt.Println("cmd error!")
		// }
	}
	params2:=[]string{"select"}
	bind2:=bind(func2,params2)
	// bind1(strings.NewReader("cat"),os.Stdout)
	fmt.Println("--------------------")
	pp:=pipe(bind1,bind2)
	pp(strings.NewReader("cat"),os.Stdout)
	//cat note|grep select
	
	{
		// pr, pw := io.Pipe()
		// defer pw.Close()
		 
		// cmd := exec.Command("cat", "test")
		// cmd.Stdout = pw
		 
		// go func() {
		//     defer pr.Close()
		//     if _, err := io.Copy(os.Stdout, pr); err != nil {
		//         fmt.Println(err)
		//     }
		// }()
		// if err := cmd.Run(); err != nil {
		//     fmt.Println(err)
		// }
	}


	time.Sleep(2*time.Second)
	fmt.Println("done!")
}


