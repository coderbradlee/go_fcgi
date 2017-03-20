export GOPATH=/root/go_fcgi
# export GODEBUG="gctrace=1"

rm -fr xx
go build -race -o xx xx
time -p ./xx
