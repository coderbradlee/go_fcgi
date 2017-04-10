package main

import (
	"fmt"
	
	"time"
)
type Fetcher interface{
	Fetch()(items []Item,next time.Time,err error)
}
func NewFetcher(domain string)Fetcher {
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
	next=time.Now().Add(time.Second)
	return
}
type Subscription interface{
	Updates()<-chan Item
	Close()error
}
type sub struct{
	fetcher Fetcher
	updates chan Item
	closing chan int
	err error
}
func (s *sub)loop() {
	//call fetch
	//send items on the updates channel
	//exit when close is called,reporting any error
	// for{
	// 	if s.closed{
	// 		close(s.updates)
	// 		return
	// 	}
	// 	items,next,err:=s.fetcher.Fetch()
	// 	if err!=nil{
	// 		s.err=err
	// 		time.Sleep(10*time.Second)
	// 		continue
	// 	}
	// 	for _,item:=range items{
	// 		s.updates<-item
	// 	}
	// 	if now:=time.Now();next.After(now){
	// 		a:=next.Sub(now)
	// 		fmt.Println("after:",a)
	// 		time.Sleep(a)
	// 	}
	// }
	
	for{
		select{
			case <-s.closing:
				// fmt.Println(cl)
				close(s.updates)
				return
			default:
				// fmt.Println("default")
		}
		// fmt.Println("after select")
		items,next,err:=s.fetcher.Fetch()
		if err!=nil{
			s.err=err
			time.Sleep(10*time.Second)
			continue
		}
		for _,item:=range items{
			// fmt.Println("item input s.updates")
			s.updates<-item
		}
		if now:=time.Now();next.After(now){
			a:=next.Sub(now)
			fmt.Println("after:",a)
			time.Sleep(a)
		}
	}
}
func (s *sub)Updates()<-chan Item {
	return s.updates
}
func (s *sub)Close()error {
	s.closing<-1
	return s.err
}
func NewSubscription(fetcher Fetcher)Subscription {
	updates:=make(chan Item)
	cl:=make(chan int)	
	s:= &sub{fetcher,updates,cl,nil}
	go s.loop()
	return s
}
type fetcherall struct{
	domains []string
}
func (f *fetcherall)Fetch()(items []Item,next time.Time,err error) {
	for _,dos:=range f.domains{
		for i:=0;i<3;i++{
			items=append(items,Item{dos,i})
		}
	}
	
	next=time.Now().Add(time.Second)
	return
}
type mergedSub struct{
	subs []Subscription
}
func (s *mergedSub)Updates()<-chan Item {
	chans:=make(chan Item)
	for{
		for i:=0;i<len(s.subs);{
			select {
				case ret:=<-s.subs[i].Updates():
					fmt.Println("ret:",ret)
					chans <-ret
					// return chans
				default:
					fmt.Println("default")
					return chans
			}
			// return s.subs[i].Updates()
		}
	}
	
	return chans
}
func (s *mergedSub)Close()error {
	for _,sub:=range s.subs{
		sub.Close()
	}
	return nil
}
func Merge(subs ...Subscription)Subscription {
	// updates:=make(chan Item)
	// var items []Item 
	// var domains []string
	// for _,sub:=range subs{
	// 	domains=append(domains,sub.fetcher.domain)
	// }
	// fa:=&fetcherall{domains}
	// cl:=make(chan int)	
	// s:= &sub{fa,updates,cl,nil}
	// go s.loop()
	s:=&mergedSub{subs}
	return s
}

func main() {
	merged:=Merge(NewSubscription(NewFetcher("xx.com")),NewSubscription(NewFetcher("xx.org")),NewSubscription(NewFetcher("xx.cn")))

	time.AfterFunc(3*time.Second,func() {
		fmt.Println("closed:",merged.Close())
	})
	for it:=range merged.Updates(){
		fmt.Println(it.Title,it.Channel)
	}
	// s:=NewSubscription(NewFetcher("xx.com"))

	// time.AfterFunc(3*time.Second,func() {
	// 	fmt.Println("closed:",s.Close())})
	// fmt.Println("weird thing happend")
	// for it:=range s.Updates(){
	// 	fmt.Println(it.Title,it.Channel)
	// }
	// fmt.Println("weird thing happend again")
	// panic("show me the stacks")
}
