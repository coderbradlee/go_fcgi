package main
import (
	"fmt"
	"runtime"
	"sync"
	"runtime/pprof"
	"os"
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
		j:=i
		fmt.Println(j)
		deferlock()
	}
}
func fib(n int)int {
	if n<2{
		return n
	}
	return fib(n-1)+fib(n-2)
}
const m1 = 0x5555555555555555
const m2 = 0x3333333333333333
const m4 = 0x0f0f0f0f0f0f0f0f
const h01 = 0x0101010101010101
func popcnt(x uint64) uint64 {
	x -= (x >> 1) & m1
	x = (x & m2) + ((x >> 2) & m2)
	x = (x + (x >> 4)) & m4
	return (x * h01) >> 56
}

func main() {
	// cpu, _ := os.Create("cpu.out")
	// defer cpu.Close()
	// pprof.StartCPUProfile(cpu)
	// defer pprof.StopCPUProfile()
	// Memory
	mem, _ := os.Create("mem.out")
	defer mem.Close()
	defer pprof.WriteHeapProfile(mem)

	test()
	fmt.Println("done")
}
