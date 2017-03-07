export GOPATH=/root/go_fcgi
export MARTINI_ENV=production
export CGO_LDFLAGS="-L/root/go_fcgi/src/mainproject/wkhtmltox"
export LD_LIBRARY_PATH="-L/root/go_fcgi/src/mainproject/wkhtmltox"
rm -fr go_fcgi
go build -o go_fcgi mainproject
./go_fcgi
