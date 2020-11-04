package LinkNodes

import "fmt"

type Node struct {
  //数据域
  Data  int
  NextNode  *Node
}

func CreateHeadNode(data int) *Node{
  var node = &Node{data , nil}
  return node
}
func AddNode(data int,node *Node) *Node{
  var newNode = &Node{data,nil}
node.NextNode = newNode
  return newNode
}
func ShowNodes(head *Node) {
  node := head
  for {
    fmt.Println(node.Data)
    if node.NextNode == nil {
      break
    } else {

    node = node.NextNode
  }
}
}

