export GOPATH=/root/go_fcgi
go build -gcflags "-N -l" src/xx/test8.go
gdb test8
set auto-load safe-path /
b test8.go:main