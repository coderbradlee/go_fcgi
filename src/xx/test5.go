package main

import (
	"fmt"
	"regexp"
	"bytes"
	"bufio"
	"os"
	"io"
	"time"
)
type Job struct{
	filename string
	results chan<- Results
}
func (j *Job)Do(reg *regexp.Regexp) {
	// fmt.Println("Do")
	// var lino int32=1111
	// j.results<-Results{j.filename,lino,"11"}
	file,err:=os.Open(j.filename)
	if err!=nil{
		fmt.Println("error:%s",err)
	}
	defer file.Close()
	reader:=bufio.NewReader(file)
	var lino int32
	for lino=1;;lino++{
		line,err:=reader.ReadBytes('\n')
		// func (b *Reader) ReadBytes(delim byte) ([]byte, error)
		line=bytes.TrimRight(line,"\r\n")
		if reg.Match(line){
			j.results<-Results{j.filename,lino,string(line)}
		}
		if err!=nil{
			if err!=io.EOF{
				fmt.Println("err:%d:%s",lino,err)
			}
			break;
		}
	}
}
type Results struct{
	filename string
	lino int32
	line string
}
func addJobs(jobs chan<- Job,filename string,results chan<- Results) {
	jobs<- Job{filename,results}
	close(jobs)
}
func doJobs(done chan<- struct{},reg *regexp.Regexp,jobs <-chan Job) {
	for job:=range jobs{
		job.Do(reg)
	}
	done<-struct{}{}
}
func waitCompletion(
	timeout int64,
	done <-chan struct{},
	results <-chan Results) {
	finish:=time.After(time.Duration(timeout))

	for i:=0;i<8;{
		select{
			case r:=<-results:
				fmt.Printf("63:%s:%d:%s\n",r.filename,r.lino,r.line)
			case <-done:
				fmt.Printf("done:%d:::::::::::\n",i)
				i++
			case <-finish:
				fmt.Printf("timeout\n")
				return
		}
	}

	for{
		select{
		case r:=<-results:
			fmt.Printf("75:%s:%d:%s\n",r.filename,r.lino,r.line)
		case <-finish:
			fmt.Printf("timeout\n")
			return
		default:
			return
		}
	}
}
func grep(reg *regexp.Regexp,filename string) {
	jobs:=make(chan Job,8)
	results:=make(chan Results,1000)
	done:=make(chan struct{},8)
	go addJobs(jobs,filename,results)
	for i:=0;i<8;i++{
		go doJobs(done,reg,jobs)
	}
	waitCompletion(100 * 1000 * 1000 * 1000,done,results)
}
func main() {
	reg:=regexp.MustCompile("flow.*")
	grep(reg,"/root/redisRenesola-cluster-debug/cache_20170224_00208.log")
}
