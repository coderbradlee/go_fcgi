package main

import (
	"fmt"
	// "sync"
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
	Close() []error
}
type concreteSub struct{
	fetchers []Fetcher
	updates chan Item
	closing chan int
	errs []error
}

func (s *concreteSub)loop() {
	for _,f:=range s.fetchers{
		go func() {
			for{
				select{
				case <-s.closing:
					fmt.Println("select returnnnnnnnnnnnnnnnnnnnnnnnnnnnnn")
					return
				default:

				}
				fmt.Println("45:",f.(*fet).domain)
				// fmt.Println("after select")
				items,next,err:=f.Fetch()
				if err!=nil{
					s.errs=append(s.errs,err)
					time.Sleep(10*time.Second)
				}
				for _,item:=range items{
					fmt.Println("item input s.updates")
					defer func() {
				        if r := recover(); r != nil {
				            err = fmt.Errorf("%v", r)
				            fmt.Printf("write: error writing %d on channel: %v\n", 1, err)
				            return
				        }
				        fmt.Printf("write: wrote %d on channel\n", 1)
				    }()
					s.updates<-item
				}
				if now:=time.Now();next.After(now){
					a:=next.Sub(now)
					fmt.Println("after:",a)
					time.Sleep(a)
				}
			}
		}()	
	}
}
func (s *concreteSub)Updates()<-chan Item {
	return s.updates
}
func (s *concreteSub)Close()[]error {
	// go func() {
	// 	<-s.closing
	// 	fmt.Println("closing")
	// 	close(s.updates)

	// }()
	go func() {
		s.closing<-1
	}()
	go func() {
		s.closing<-1
	}()
	go func() {
		s.closing<-1
	}()
	fmt.Println("closing/////////////////////////////")
	return s.errs
}
func NewSubscription(fetcher Fetcher)Subscription {
	updates:=make(chan Item)
	cl:=make(chan int)	
	fet:=[]Fetcher{fetcher}
	s:= &concreteSub{fet,updates,cl,nil}
	s.loop()
	return s
}

func Merge(subs ...Subscription)Subscription {
	updates:=make(chan Item)
	cl:=make(chan int)	
	fets:=[]Fetcher{}
	merged:= &concreteSub{fets,updates,cl,nil}

	for _,s:=range subs{
		convert:=s.(*concreteSub)
		merged.fetchers=append(merged.fetchers,convert.fetchers...)
		s.Close()
	}
	// for _,s:=range merged.fetchers{
	// 	fmt.Println("111:",s.(*fet).domain)
	// }
	
	merged.loop()
	return merged
}

func main() {
	var domains =[]string{"xx.com","yy.com","zz.com"}
	var subs []Subscription
	for _,domain:=range domains{
		subs=append(subs,NewSubscription(NewFetcher(domain)))
	}
	mainFeed:=Merge(subs...)
	updates:=mainFeed.Updates()
	time.AfterFunc(5*time.Second,func(){mainFeed.Close()})
	for{
		select{
		case it:=<-updates:
			fmt.Println(it.Title,it.Channel)
		// case <-time.After(5*time.Second):
		// 	fmt.Println("mainfeed close//////////////////////////////////")
		// 	mainFeed.Close()
		}
	}
}
