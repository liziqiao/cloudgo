# cloudgo
-----------------------
##martini

**这次作业我选择了使用martini框架来完成，下面是go语言中文网对于这个框架的介绍：**

Martini 是一个非常新的 Go 语言的 Web 框架，使用 Go 的 net/http 接口开发，类似 Sinatra 或者 Flask 之类的框架，你可使用自己的 DB 层、会话管理和模板。

特性：
使用非常简单
无侵入设计
可与其他 Go 的包配合工作
超棒的路径匹配和路由
模块化设计，可轻松添加工具
大量很好的处理器和中间件
很棒的开箱即用特性
完全兼容 http.HandlerFunc 接口.

使用go get命令即可获取
``` 
go get github.com/go-martini/martini
```
这次的作业还是用到了上次的pflag包来解析命令行：
``` go
import (
  "os"
  server "github.com/liziqiao/cloudgo/service"
  flag "github.com/spf13/pflag"
)
```
服务端的代码中使用了框架之后处理GET的方法就显得非常简单了
```
m := martini.Classic()
m.Get("/", func(params martini.Params) string {
     return "helloworld!"
})
m.RunOnAddr(":"+port)
```

运行main.go,使用参数-p8000，然后在命令行输入curl命令
输入curl命令的情况：
```
liziqiao@PC-201508262204:~$ curl -v http://localhost:8000/
* Hostname was NOT found in DNS cache
*   Trying 127.0.0.1...
* Connected to localhost (127.0.0.1) port 8000 (#0)
> GET / HTTP/1.1
> User-Agent: curl/7.35.0
> Host: localhost:8000
> Accept: */*
>
< HTTP/1.1 200 OK
< Date: Tue, 14 Nov 2017 16:13:43 GMT
< Content-Length: 11
< Content-Type: text/plain; charset=utf-8
<
* Connection #0 to host localhost left intact
helloworld!
```
服务端的情况：
```
C:\Users\Administrator\Desktop\gg\cloudgo>go run main.go -p 8000
[martini] listening on :8000 (development)
[martini] Started GET / for 127.0.0.1:35138
[martini] Completed 200 OK in 2.0047ms
```
用浏览器打开http://localhost:8000/可以看到helloworld!

进行压力测试，这里安装了windows版本的apache，测试结果如下，括号内的是自己家的关键参数的解释：
```
C:\Users\Administrator\Downloads\httpd-2.4.29-o102m-x64-vc14\Apache24\bin
λ ab -n 1000 http://localhost:8000/
This is ApacheBench, Version 2.3 <$Revision: 1807734 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:
Server Hostname:        localhost
Server Port:            8000

Document Path:          /（这个代表请求的资源）
Document Length:        11 bytes（文档返回的长度）

Concurrency Level:      1（并发个数）
Time taken for tests:   1.681 seconds（总的请求时间）
Complete requests:      1000（总的请求个数）
Failed requests:        0（失败的请求个数）
Total transferred:      128000 bytes
HTML transferred:       11000 bytes
Requests per second:    594.82 [#/sec] (mean)（平均每秒的请求数）
Time per request:       1.681 [ms] (mean)（平均每个请求消耗的时间）
Time per request:       1.681 [ms] (mean, across all concurrent requests)
Transfer rate:          74.35 [Kbytes/sec] received（传输速率）

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.4      0       2
Processing:     0    1   2.1      1      10
Waiting:        0    1   2.0      1      10
Total:          0    2   2.1      1      10

Percentage of the requests served within a certain time (ms)
  50%      1（相应百分比的请求的完成时间）
  66%      1
  75%      1
  80%      2
  90%      4
  95%      8
  98%      9
  99%      9
 100%     10 (longest request)
```



