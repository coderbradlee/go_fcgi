export GOPATH=/root/go_fcgi
export GOMAXPROCS=1
rm -fr go_fcgi
go build -o go_fcgi mainproject
./go_fcgi
