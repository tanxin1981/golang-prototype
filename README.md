beego-prototype
===============

Prototype using beego

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

