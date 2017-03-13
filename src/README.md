2017.3.7
新环境安装：
1、yum install go
cp src/wkhtmltox/libwkhtmltox.so.0.12.3 /usr/lib
ln -s /usr/lib/libwkhtmltox.so.0.12.3 /usr/lib/libwkhtmltox.so.0
ln -s /usr/lib/libwkhtmltox.so.0.12.3 /usr/lib/libwkhtmltox.so
或
yum install wkhtmltopdf* 不能正常工作

curl -X POST http://127.0.0.1:9888/pdf -d '{"src":"/root/Newshop_International/WEB-INF/quotation_pdf.html","dst":"test.pdf"}'

curl -X POST http://127.0.0.1:9888/po/deliver_goods -d '{"xx":"xxxx"}'




curl -X POST http://172.18.100.85/pdf -d '{"src":"/root/Newshop_International/WEB-INF/quotation_pdf.html","dst":"test.pdf"}'

ab -n 1000000 -k -r -c 100 http://127.0.0.1/pdf
./weighttp -n 500000 -c 100 -t 2 -k 127.0.0.1/pdf

./weighttp -n 500000 -c 100 -t 2 -k 127.0.0.1:9888/redis