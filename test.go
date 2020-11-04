package main
//王嗣萱2018110601014
import "fmt"
var(
    insertNumber ,a    int
)
func main(){
    fmt.Scanf("%d",&insertNumber)
    for a:=0;a < insertNumber ; a++{
    	var b int = insertNumber-a
        fmt.Printf("%d + %d = %d \n",a,b,insertNumber)
    }
}