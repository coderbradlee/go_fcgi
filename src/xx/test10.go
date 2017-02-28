package main

import (
	"fmt"
	"regexp"
	"os"
	"bufio"
	"runtime"
	"io"
	"sync"
)
type PageMap struct{
	count map[string]int
	mutex *sync.RWMutex
}
var workers=runtime.NumCPU()
func main() {
	// test :=New()
	// test.Insert("1",2)
	// if data,found:=test.Find("1");found{
	// 	fmt.Println(data)
	// }
	filename:="/root/redisRenesola-cluster-debug/logs/cache_20170227_00209.log"
	runtime.GOMAXPROCS(runtime.NumCPU())
	lines := make(chan string, workers*4)
    done := make(chan struct{}, workers)
    pageMap := PageMap{make(map[string]int,new(sync.RWMutex))}
    go readLines(filename, lines)
    processLines(done, pageMap, lines)
    waitUntil(done)
    pageMap.showResults()
}
func readLines(filename string, lines chan<- string) {
    file, err := os.Open(filename)
    if err != nil {
        fmt.Println("failed to open the file:", err)
    }
    defer file.Close()
    reader := bufio.NewReader(file)
    for {
        line, err := reader.ReadString('\n')
        if line != "" {
            lines <- line
        }
        if err != nil {
            if err != io.EOF {
                fmt.Println("failed to finish reading the file:", err)
            }
            break
        }
    }
    close(lines)
}
func (p *PageMap)Increment(page string) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.count[page]++
}
func processLines(done chan<- struct{}, pageMap PageMap,
    lines <-chan string) {
    getRx := regexp.MustCompile(`GET /flowNo/.?`)
    
    for i := 0; i < workers; i++ {
        go func() {
            for line := range lines {
                if matches := getRx.FindStringSubmatch(line);
                    matches != nil {
                    pageMap.Increment(matches[0])
                }
            }
            done <- struct{}{}
        }()
    }
}

func waitUntil(done <-chan struct{}) {
    for i := 0; i < workers; i++ {
        <-done
    }
}

func (p *PageMap)showResults() {
	p.mutex.RLock()
	defer p.mutex.RUnlock()
    for page, count := range p.count {
        fmt.Printf("%s %8d\n",page, count )
    }
}
