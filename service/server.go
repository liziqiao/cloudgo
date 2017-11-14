package service

import (
  "github.com/go-martini/martini"
)

//使用martini框架
func NewServer(port string) {
  m := martini.Classic() //创建一个典型的martini实例
  //处理GET请求，第二个参数是处理函数
  m.Get("/", func(params martini.Params) string {
      return "helloworld!"
  })
  m.RunOnAddr(":"+port)//运行服务器
}
