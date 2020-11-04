package main

import (
  "fmt"
  "math"
  "net/http"
  "strconv"
)
func main() {
  http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
    //学习需要，简单解决跨域问题
    writer.Header().Set("Access-Control-Allow-Origin", "*")
    //获取url中的数据
    number1 := request.FormValue("num1")
    number2 := request.FormValue("num2")
    //转换string成float64格式
    num1, err := strconv.ParseFloat(number1,64)
    num2, err := strconv.ParseFloat(number2,64)
    var result float64
    if err != nil {
      http.Error(writer, err.Error(), http.StatusInternalServerError)
      return
    }
    operator := request.FormValue("operator")
    switch operator {
    case "add":result = add(num1,num2)
    case "minus":result = minus(num1,num2)
    case "multiply":result = multiply(num1,num2)
    case "divide":result = divide(num1,num2)
    case "remainder":result = remainder(num1,num2)
    case "root":result = root(num1)
    case "square":result = square(num1)
    case "reciprocal":result = reciprocal(num1)
    case "cos":result = cos(num1)

    }
    //fmt.Fprintln(writer,fmt.Sprintf("数1：%f,数2:%f,符号:%s",num1,num2,operator))
    //返回json数据
    fmt.Fprintln(writer,fmt.Sprintf("%f",result))
    //type Result struct {
    //  result  float64
    //}
    //result00 := Result{result}
    //js, err := json.Marshal(result00)
    //if err != nil {
    //  http.Error(writer, err.Error(), http.StatusInternalServerError)
    //  return
    //}
    //writer.Write(js)
  })
  fmt.Println(http.ListenAndServe(":8080",nil))
}
func add(num1 float64,num2 float64)(value float64){
  value = num1 + num2
  return
}
func minus(num1,num2 float64)(value float64){
  value = num1 - num2
  return
}
func multiply(num1,num2 float64)(value float64){
  value = num1 * num2
  return
}
func divide(num1,num2 float64)(value float64){
  value = num1 / num2
  return
}
func remainder(num1,num2 float64)(value float64){
  //求余运算发生在整数上
  x := int(num1)
  y := int(num2)
  z := x % y
  w :=strconv.Itoa(z)
  value,err := strconv.ParseFloat(w,64)
  fmt.Println(err)
  return
}

func root(num1 float64)(value float64){
  value = math.Sqrt(num1)
  return
}
func square(num1 float64)(value float64){
  value = num1 * num1
  return
}
func reciprocal(num1 float64)(value float64){
  value = 1 / num1
  return
}
func cos(num1 float64) (value float64){
  value = math.Cos(num1)
  return
}
