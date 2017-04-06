package main
import(
	"crypto/rand"
	"crypto/tls"
	"io"
	"fmt"
	"net"
	"time"
)
func check(err error) {
	if err!=nil{
		fmt.Println(err)
	}
}
func handleClient(conn net.Conn) {
	defer conn.Close()
	buf:=make([]byte,512)
	for{
		fmt.Println("waiting..")
		n,err:=conn.Read(buf)
		if err!=nil{
			check(err)
			if err!=io.EOF{
				check(err)
			}
			break
		}
		fmt.Println("echo",string(buf[:n]))
		n,err=conn.Write(buf[:n])
		fmt.Println("write:",n)
		if err!=nil{
			check(err)
			break
		}
	}
	fmt.Println("over")
}
func main() {
	cert,err:=tls.LoadX509KeyPair("/root/cert/server.crt","/root/cert/server.key")
	check(err)
	config:=tls.Config{Certificates:[]tls.Certificate{cert}}
	config.Time=time.Now
	config.Rand=rand.Reader
	service:=":8000"
	listener,err:=tls.Listen("tcp",service,&config)
	check(err)
	fmt.Println("listener...")
	for{
		conn,err:=listener.Accept()
		if err!=nil{
			check(err)
			break
		}
		fmt.Println("server accept:",conn.RemoteAddr())
		go handleClient(conn)
	}
}