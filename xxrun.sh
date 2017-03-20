export GOPATH=/root/go_fcgi
# export GODEBUG="gctrace=1"
export GOBIN=/root/go_fcgi
rm -fr xx
go build -race -o xx xx
time -p ./xx
