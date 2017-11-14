# cloudgo

###martini
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

由于安装压力测试工具还有点问题，压力测试还没有完成

