package main

import (
	"dataStructure/linkedList"
	"fmt"
)

func main() {

	list := new(linkedList.List)

	n := new(linkedList.Node)

	list.PushBack(n)

	n2 := new(linkedList.Node)
	list.PushBack(n2)

	front := list.Front()

	for front != nil {
		fmt.Println("Value: ", front)
		front = front
		break
	}

	fmt.Println(list)
	fmt.Println("__ here __")

	//node2 := &Node{nil, 30}
	//node3 := &Node{nil, 40}
	//node4 := &Node{nil, 50}

}
