package main
import (
	"fmt"
	"runtime"
	"sync"
	"runtime/pprof"
)

var lock sync.Mutex

func init() {
	// runtime.NumCPU()
	runtime.GOMAXPROCS(runtime.NumCPU())
}
func deferlock() {
	lock.Lock()
	defer lock.Unlock()	
}
func test() {
	for i:=0;i<1000000;i++{
		deferlock()
	}
}
func main() {
	cpu, _ := os.Create("cpu.out")
	defer cpu.Close()
	pprof.StartCPUProfile(cpu)
	defer pprof.StopCPUProfile()
	// Memory
	// mem, _ := os.Create("mem.out")
	// defer mem.Close()
	// defer pprof.WriteHeapProfile(mem)

	test()
	fmt.Println("done")
}
