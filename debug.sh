export GOPATH=/root/go_fcgi
export GOBIN=/root/go_fcgi
go build -gcflags "-N -l" -race xx -o xx
gdb xx