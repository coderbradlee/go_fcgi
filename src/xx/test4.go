package main

import (
	"fmt"
	"regexp"
	"bytes"
	"bufio"
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
	for lino:=1;;lino++ {
		line,err:=reader.ReadBytes('\n')
		// func (b *Reader) ReadBytes(delim byte) ([]byte, error)
		line=bytes.TrimRight(line,'\r\n')
		if reg.Match(line){
			j.results<-Results{j.filename,lino,line}
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
func waitCompletion(done <-chan struct{},results chan<- Results) {
	for i:=0;i<8;i++{
		<-done
	}
	close(results)
}
func processResults(results <-chan Results) {
	for i:=range results{
		fmt.Printf("%s:%d:%s",i.filename,i.lino,i.line)
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
	go waitCompletion(done,results)
	processResults(results)
}
func main() {
	reg,_:=regexp.Compile("*flow*")
	grep(reg,"/root/redisRenesola-cluster-debug/cache_20170224_00208.log")
}
