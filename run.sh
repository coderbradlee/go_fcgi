export GOPATH=/root/go_fcgi
export GOMAXPROCS=1
export MARTINI_ENV=production
export GODEBUG="memprofilerate=1"
rm -fr go_fcgi
go build -o go_fcgi mainproject
./go_fcgi
