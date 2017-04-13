export GOPATH=/root/go_fcgi
# export GODEBUG="gctrace=1"
export GOBIN=/root/go_fcgi
rm -fr testpprof
go build -race -o testpprof testpprof
# time -p ./testpprof
