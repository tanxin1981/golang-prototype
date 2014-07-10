golang-prototype
===============

Prototype using beego

```
setup:
$ sudo apt-get install golang
$ export GOPATH=<path_of_your_choice>
$ go get github.com/astaxie/beego
$ go get github.com/astaxie/beego/orm
$ go get github.com/go-sql-driver/mysql
$ cd <GOPATH>/src
$ git clone git@github.com:zhengyang4k/golang-prototype.git
$ cd golang-prototype
$ go run main.go

db settings: <username>:<password>@/<dbname>?charset=utf8
main.go:
...
func init() {
	// register database
	orm.RegisterDataBase("default", "mysql", "root@/sgfas?charset=utf8", 30)
}
...
```

```
test urls:
simple: http://localhost:8080/v1/object
complex(with db connection): http://localhost:8080/v1/catalog_simple/10002
```


```
Benchmark: ab -n 10000 -c 70 <url>

Simple string:

Total transferred:      3580000 bytes
HTML transferred:       2080000 bytes
Requests per second:    22556.71 [#/sec] (mean)
Time per request:       3.103 [ms] (mean)
Time per request:       0.044 [ms] (mean, across all concurrent requests)
Transfer rate:          7886.04 [Kbytes/sec] received


DB primary key look up with 1 join:

Total transferred:      2610000 bytes
HTML transferred:       1110000 bytes
Requests per second:    13376.67 [#/sec] (mean)
Time per request:       5.233 [ms] (mean)
Time per request:       0.075 [ms] (mean, across all concurrent requests)
Transfer rate:          3409.48 [Kbytes/sec] received
```

