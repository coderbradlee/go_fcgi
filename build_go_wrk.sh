export GOPATH=/root/go_fcgi
export GOBIN=/root/go_fcgi
# go get -v github.com/uber/go-torch
# go-torch -u http://localhost:9888/redis -t 30
rm -fr go-wrk
go build -o go-wrk go-wrk