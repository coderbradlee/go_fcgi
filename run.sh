export GOPATH=/root/go_fcgi
export GOMAXPROCS=1
go build -o go_fcgi mainproject go-redis-cluster
./go_fcgi
