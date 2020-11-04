package main

import (
  "fmt"
  "log"
  "net/http"
)
//定义价格
type dollars float32

func (d dollars) String() string{ return fmt.Sprintf("$%.2f",d)}
//通过map模型化数据库 键值为 string dollars
type database map[string]dollars

//遍历数据库database来打印在网页
func (db database) ServeHTTP(v http.ResponseWriter, req *http.Request){
  for item ,price := range  db {
    fmt.Fprintf(v, "%s:%s\n",item,price)
  }
}
//主程序，打开本地8000端口
func main()  {
  db := database{"shoes":50,"socks":5}
  log.Fatal(http.ListenAndServe("localhost:8000",db))
}
