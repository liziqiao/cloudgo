package main

import (
  "os"
  server "github.com/liziqiao/cloudgo/service"
  flag "github.com/spf13/pflag"
)

const (
  //设置默认端口为8080
  PORT string = "8080"
)

func main() {
  port := os.Getenv("PORT")//获取去环境变量，看是否存在PORT
  if len(port) == 0 {
      port = PORT
  }
  //解析命令行
  pPort := flag.StringP("port", "p", PORT, "listening PORT")
  flag.Parse()
  if len(*pPort) != 0 {
      port = *pPort
  }
  //运行服务端
  server.NewServer(port)
}
