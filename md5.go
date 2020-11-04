package main

import (
  "crypto/md5"
  "fmt"
)
//MD5
func MyMD5(message string){
  data:=[]byte(message)
  fmt.Println(data)
  s:=fmt.Sprintf("%x",md5.Sum(data))
  fmt.Println(md5.Sum(data))
  fmt.Println(s)

}
func main(){
  MyMD5("hello world")
}
