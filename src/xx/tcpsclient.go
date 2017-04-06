package main
import(
	"crypto/tls"
	"io"
	"fmt"
	"crypto/x509"
)
const rootPEM=`-----BEGIN CERTIFICATE-----
MIID9TCCAt2gAwIBAgIJAMPGGv6mYYp4MA0GCSqGSIb3DQEBCwUAMIGQMQswCQYD
VQQGEwJjbjERMA8GA1UECAwIc2hhbmdoYWkxETAPBgNVBAcMCHNoYW5naGFpMRww
GgYDVQQKDBNEZWZhdWx0IENvbXBhbnkgTHRkMQswCQYDVQQLDAJpdDEPMA0GA1UE
AwwGeHguY29tMR8wHQYJKoZIhvcNAQkBFhBsenhtMTYwQHNpbmEuY29tMB4XDTE3
MDQwNjA3MTU1NVoXDTE3MDUwNjA3MTU1NVowgZAxCzAJBgNVBAYTAmNuMREwDwYD
VQQIDAhzaGFuZ2hhaTERMA8GA1UEBwwIc2hhbmdoYWkxHDAaBgNVBAoME0RlZmF1
bHQgQ29tcGFueSBMdGQxCzAJBgNVBAsMAml0MQ8wDQYDVQQDDAZ4eC5jb20xHzAd
BgkqhkiG9w0BCQEWEGx6eG0xNjBAc2luYS5jb20wggEiMA0GCSqGSIb3DQEBAQUA
A4IBDwAwggEKAoIBAQDgeHLTycwbp5AJW0vETapoFArkwtFQVWI0W3p3BZECxcPL
p/E+AXmP5oYhy5tLZ1SYfOk7Ntc4gunaasoMkZWJhcejrhVSk/tMxxEXoAmbuapZ
gpswk1jUBdJIoQd3XFahjZuaUQNgXM3F5vU5Ne5x9jm39C/1i/OlbxeuK0LKiVSz
OnBCS2QFToSfVi4QypZIJszknsgDxoJQcPWqrs8HCIUDYUDRzzMfr16zoK8nfStf
NUVXQfp8zLQEiVudEAdGffUGKu/CuuuaPniKdkbwNhPWlUDtZSv6HYI5MYZ0XDOs
oF7s+C96JkvxVYlidkU8gGI4ALc0lJteVmA0QQjXAgMBAAGjUDBOMB0GA1UdDgQW
BBSPbD9L/yO/04ZUrBiWNZRt5sgpGzAfBgNVHSMEGDAWgBSPbD9L/yO/04ZUrBiW
NZRt5sgpGzAMBgNVHRMEBTADAQH/MA0GCSqGSIb3DQEBCwUAA4IBAQCkCwMmua6T
CM1YJdZczV8XxscKrcyYAgl7H50WXqpH2laF6iqtMZhXNoLSTr7yCpwrcVmsD3sn
wwyVJBMfOz/32TnqD4U1CT2yb9NMLFG6ApIQNx40B2aliJ/6KiaEz0z1ja7tn36N
ApPfPebG3R/MOUjrR3asO2vAcUErtzmtA6sGRrWL/7qYXRF27WhJ8YycUrFPCAFq
OvZcLbKD3g8uFyffbMU8YTrwisqQK3Miqt7oMNL9a8g0V7VZRuFEAitomo4k9PUf
lFnUFDcaGjovPzkCIylpudUrljIS2DDBpJ5od3/e2cZJG8Qv+nFLgg4GY4k4fqTA
+7YhGmtu6mfp
-----END CERTIFICATE-----`
func check(err error) {
	if err!=nil{
		fmt.Println(err)
	}
}

func main() {
	// conn,err:=tls.Dial("tcp","xx.com:8000",&tls.Config{
 //    InsecureSkipVerify: true,ServerName:"xx.com",})
	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM([]byte(rootPEM))
	if !ok {
	    panic("failed to parse root certificate")
	}
	conn,err:=tls.Dial("tcp","xx.com:8000",&tls.Config{RootCAs: roots})
	check(err)
	defer conn.Close()
	fmt.Println("conn to:",conn.RemoteAddr())
	state:=conn.ConnectionState()
	fmt.Println("handshake:",state.HandshakeComplete)
	fmt.Println("mutual:",state.NegotiatedProtocolIsMutual)
	message:="hello\n"
	n,err:=io.WriteString(conn,message)
	check(err)
	fmt.Println("write:",message)
	reply:=make([]byte,256)
	n,err=conn.Read(reply)
	fmt.Println("read:",string(reply[:n]))
	fmt.Println("exiting")
}