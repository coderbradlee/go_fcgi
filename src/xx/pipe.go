package main

//#go:generate ls -l
import (
	"fmt"
	
)

func pipe(app1 func(in io.Rader,out io.Writer),app2 func (in io.Rader,out io.Writer))func(in io.Rader,out io.Writer) {
	return func(in io.Rader,out io.Writer) {
		pr,pw:=io.Pipe()
		defer pw.Close()
		go func() {
			defer pr.Close()
			app2(pr,out)
		}()
		app1(in,pw)
	}
}
func main() {
	// testreflect()
	testreflect1()
	
}
