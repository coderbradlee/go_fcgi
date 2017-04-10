package main

import (
	"fmt"
	
	"net"
	"net/http"
	"reflect"
)

type Fetcher interface{
	Fetch()(items []Item,next time.Time,err error)
}
func Fetch(domain string)Fetcher{} {
	
}
type Subscription interface{
	Updates()<-chan Item
	Close()error
}
func Subscribe(fetcher Fetcher)Subscription{...} {
	
}
func Merge(subs ...Subscription)Subscription{...} {
	
}
type Item struct{
	Channel string
	Title string
}
func main() {
	merged:=Merge(Subscribe(Fetch("xx.com")),Subscribe(Fetch("xx.org")),Subscribe(Fetch("xx.cn")))
	time.AfterFunc(3*time.Second,func() {
		fmt.Println("closed:",merged.Close())
	})
	for it:=range merged.Updates(){
		fmt.Println(it.Channel,it.Title)
	}
	panic("show me the stacks")
}
