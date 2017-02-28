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
    pageMap := PageMap{make(map[string]int),new(sync.RWMutex)}
    go readLines(filename, lines)
    processLines(done, pageMap, lines)
    waitUntil(done)
    pageMap.showResults()
}

