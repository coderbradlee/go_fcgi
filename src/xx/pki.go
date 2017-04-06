package main
import (
	"fmt"
	"crypto/sha1"
	"crypto/md5"
	"io"
	"os"
)

func main() {
	file:="123.txt"
	{
		infile,inerr:=os.Open(file)
		defer infile.Close()
		if inerr!=nil{
			fmt.Println(inerr)
		}else{
			md5h:=md5.New()
			io.Copy(md5h,infile)
			fmt.Printf("%x\n",md5h.Sum([]byte("")))
		}
	}
	{
		infile,inerr:=os.Open(file)
		defer infile.Close()
		if inerr!=nil{
			fmt.Println(inerr)
		}else{
			sha1h:=sha1.New()
			io.Copy(sha1h,infile)
			fmt.Printf("%x\n",sha1h.Sum([]byte("")))
		}
	}

	// test:="Hi,pandaman!"
	// md5inst:=md5.New()
	// md5inst.Write([]byte(test))
	// ret:=md5inst.Sum([]byte(""))
	// fmt.Printf("%x\n",ret)

	// sha1inst:=sha1.New()
	// sha1inst.Write([]byte(test))
	// ret=sha1inst.Sum([]byte(""))
	// fmt.Printf("%x\n",ret)
}
