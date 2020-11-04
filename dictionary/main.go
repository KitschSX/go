package main

import (
  "bufio"
  "fmt"
  "io"
  "net/http"
  "os"
)
func main() {
  file,err := os.Open("dictionary/words.txt")
  var i int = 1
  var letters [5000] string
  var attributes [5000] string
  var translates [5000] string

  if err!=nil{
    panic(err)
  }
  rd := bufio.NewReader(file)
  for{
    letter,err := rd.ReadString(' ')
    attribute,err := rd.ReadString('.')
    translate,err := rd.ReadString('\n')
    if err!=nil || io.EOF==err{
      break
    }else{
      letters[i] = letter
      attributes[i] = attribute
      translates[i] = translate
      //fmt.Printf(letter)
      //fmt.Printf(attribute)
      fmt.Printf(translate)
      i++
    }
  }
  //http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
  //  //跨域问题
  //  //writer.Header().Set("Access-Control-Allow-Origin", "*")
  //  //获取url中的数据
  //  //letter := request.FormValue("letter")
  //  fmt.Fprintln(writer,fmt.Sprintf("hello"))
  //})
  fmt.Println(http.ListenAndServe(":8080",nil))
}
