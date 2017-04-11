package main

import (
	"fmt"
	"sync"
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
	// chans:=make(chan Item)
	// go func() {
	// 	for{
	// 		for i:=0;i<len(s.subs);{
	// 			select {
	// 				case ret:=<-s.subs[i].Updates():
	// 					fmt.Println("ret:",ret)
	// 					chans <-ret
	// 				default:
	// 					fmt.Println("default")
	// 			}
	// 			// return s.subs[i].Updates()
	// 		}
	// 	}
	// }()
	// return chans
	leng:=len(s.subs)
	fmt.Println(leng)
	chans:=make(chan Item)
	var wg sync.WaitGroup
    // wg.Add(leng)
    for i, sub := range s.subs {
    	wg.Add(1)
        go func(i int, sub <-chan Item) {
            for s := range sub {
                chans <- s
                fmt.Println(s.Title,s.Channel)
            }
            defer wg.Done()
        }(i, sub.Updates())
    }
    go func() {
    	for it:=range chans{
			fmt.Println(it.Title,it.Channel)
		}
    }()
    wg.Wait()
    
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
	var domains []string=[]string{"xx.com","yy.com","zz.com"}
	var subs []Subscription
	for _,domain:=range domains{
		subs=append(subs,NewSubscription(NewFetcher(domain)))
	}
	mainFeed:=Merge(subs...)
	updates:=mainFeed.Updates()
	for{
		select{
		case it:=<-updates:
			fmt.Println(it.Title,it.Channel)
		case <-time.After(10*time.Second):
			mainFeed.Close()
		}
	}
}
