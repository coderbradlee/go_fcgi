export GOPATH=/root/go_fcgi
export GOBIN=/root/go_fcgi
export MARTINI_ENV=production

rm -fr go_fcgi
go build -o go_fcgi mainproject
./go_fcgi
#go install mainproject
