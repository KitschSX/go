package main

import (
  "crypto/sha256"
  "fmt"
)

func MySha256(message string){
  data:=[]byte(message)
  fmt.Println(data)
  s:=fmt.Sprintf("%x",sha256.Sum256(data))
  fmt.Println(sha256.Sum256(data))
  fmt.Println(s)
}
func main(){
  MySha256("hello world")
}
