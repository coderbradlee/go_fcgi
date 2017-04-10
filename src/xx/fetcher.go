package main

import (
	"fmt"
	
	"time"
)
type Fetcher interface{
	Fetch()(items []Item,next time.Time,err error)
}
func Fetch(domain string)Fetcher {
	return &fet{domain}
}
type fet struct{
	domain string
}
type Item struct{
	Title string
	Channel int
}
func (f *fet)Fetch()(items []Item,next time.Time,err error) {
	for i:=0;i<3;i++{
		items=append(items,Item{f.domain,i})
	}
	next=time.Now()
	return
}
type Subscription interface{
	Updates()<-chan Item
	Close()error
}
type sub struct{
	fetcher Fetcher
	updates chan Item
	closed bool
	err error
}
func (s *sub)loop() {
	//call fetch
	//send items on the updates channel
	//exit when close is called,reporting any error
	for{
		if s.closed{
			close(s.updates)
			return
		}
		items,next,err:=s.fetcher.Fetch()
		if err!=nil{
			s.err=err
			time.Sleep(10*time.Second)
			continue
		}
		for _,item:=range items{
			s.updates<-item
		}
		if now:=time.Now();next.After(now){
			time.Sleep(next.Sub(now))
		}
	}
}
func (s *sub)Updates()<-chan Item {
	return s.updates
}
func (s *sub)Close()error {
	s.closed=true
	return s.err
}
func Subscribe(fetcher Fetcher)Subscription {
	updates:=make(chan Item)
	s:= &sub{fetcher,updates,false,nil}
	go s.loop()
	return s
}
// func Merge(subs ...Subscription)Subscription {
	
// }

func main() {
	// merged:=Merge(Subscribe(Fetch("xx.com")),Subscribe(Fetch("xx.org")),Subscribe(Fetch("xx.cn")))
	// time.AfterFunc(3*time.Second,func() {
	// 	fmt.Println("closed:",merged.Close())
	// })
	// for it:=range merged.Updates(){
	// 	fmt.Println(it.Title,it.Channel)
	// }
	s:=Subscribe(Fetch("xx.com"))
	for it:=range s.Updates(){
		fmt.Println(it.Title,it.Channel)
	}
	panic("show me the stacks")
}
