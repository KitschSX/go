package main

import (
  "../gogo/LinkNodes"
  "fmt"
)

func main(){
  fmt.Println("hello,world")
  //调用，创建头节点
  head:=LinkNodes.CreateHeadNode(1)
  //添加新节点
  node:= LinkNodes.AddNode(2,head)
  node = LinkNodes.AddNode(3,node)
  node = LinkNodes.AddNode(4,node)
  fmt.Println(head)
  fmt.Println(head.NextNode)
  fmt.Println(head.NextNode.NextNode.NextNode)
  fmt.Println(head.NextNode.NextNode.NextNode.NextNode)
  LinkNodes.ShowNodes(head)
}
