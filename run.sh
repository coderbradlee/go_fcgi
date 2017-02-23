export GOPATH=/root/go_fcgi
export GOMAXPROCS=1
export MARTINI_ENV=production
rm -fr go_fcgi
go build -o go_fcgi mainproject
go build -o xx xx
./go_fcgi
