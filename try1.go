package main
import "fmt"
type Student struct {
  id int
  name string
  age int
  sex byte
}
func main() {
	stu:=[]Student{
	  Student{
	    id:0,name:"张三",age:15,sex:'m',
    },
    Student{
      id:1,name:"李四",age:15,sex:'m',
    },Student{
      id:2,name:"王五",age:16,sex:'m',
    },
  }
	fmt.Println(stu)
  m:=map[int]Student{1:Student{
    id:0,name:"张三",age:15,sex:'m',
  }}
  fmt.Println(m)
}
