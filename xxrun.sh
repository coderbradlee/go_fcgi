export GOPATH=/root/go_fcgi
export MARTINI_ENV=production

rm -fr xx
go build -o xx xx
time -p ./xx
