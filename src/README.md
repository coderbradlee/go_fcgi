2017.3.4 
tag allerrorchan2017.3.4
全部改为了chan传递参数的模式，但在最后的时候insert时，chan<-nil阻塞了返回值
无法解决，尝试改为errgroup.Group模式
curl -X POST http://127.0.0.1:9888/po/deliver_goods -d '{"xx":"xxxx"}'




curl -X GET http://172.18.100.85/pdf
ab -n 1000000 -k -r -c 100 http://127.0.0.1/pdf
./weighttp -n 500000 -c 100 -t 2 -k 127.0.0.1/pdf

./weighttp -n 500000 -c 100 -t 2 -k 127.0.0.1:9888/redis