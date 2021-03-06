package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func (node Node) Print() {
	fmt.Print(node.Value, " ")
}

func CreateNode(value int) *Node {
	return &Node{Value: value}
}

func (node *Node) SetValue(value int) {
	node.Value = value
}

func (node *Node) PostOrder() {
	node.TraverseFunc(func(n *Node) {
		n.Print()
	})
	fmt.Println()
}
func (node *Node) TraverseFunc(f func(node *Node)) {
	if node == nil {
		return
	}
	node.Left.TraverseFunc(f)
	node.Right.TraverseFunc(f)
	f(node)
}

//利用channel遍历
func (node *Node) TraverseWithChannel() chan *Node {
	out := make(chan *Node)
	go func() {
		node.TraverseFunc(func(node *Node) {
			out <- node
		})
		close(out)
	}()
	return out
}
