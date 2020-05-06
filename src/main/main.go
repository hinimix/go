package main

import (
	"fmt"
	"math"
	"math/rand"
)

// 二叉查询树，左边小于根节点，右边大于根节点

// 链表
type Node struct {
	data  int
	left  *Node
	right *Node
}

// 链式存储
func insertListBinaryTree(btree *Node, value int) {
	var tmpNode Node
	tmpNode.data = value
	if (*btree).data == 0 {
		(*btree).data = value
	} else if value < (*btree).data {
		if (*btree).left == nil {
			(*btree).left = &tmpNode
		} else {
			insertListBinaryTree((*btree).left, value)
		}
	} else if value > (*btree).data {
		if (*btree).right == nil {
			(*btree).right = &tmpNode
		} else {
			insertListBinaryTree((*btree).right, value)
		}
	} else {
		fmt.Println("can not insert")
	}
}
func generateListBinaryTree(btree *Node) {
	number := 20
	for i := 0; i <= number; i++ {
		var value = rand.Intn(100)
		insertListBinaryTree(btree, value)
	}
}

// 前序遍历链式, 打印顺序：当前节点，左子树，右子树
func preOrderSearchList(btree *Node) {
	if (*btree).data != 0 {
		fmt.Println((*btree).data)
	} else {
		fmt.Println("no root")
	}

	if (*btree).left != nil {
		preOrderSearchList((*btree).left)
	}
	if (*btree).right != nil {
		preOrderSearchList((*btree).right)
	}
}

// 中序遍历链式， 打印顺序：左子树，当前节点，右子树
func inOrderSearchList(btree *Node) {
	if (*btree).data !=0 {
		if (*btree).left != nil {
			inOrderSearchList((*btree).left)
		}
		fmt.Println((*btree).data)
		if(*btree).right !=nil {
			inOrderSearchList((*btree).right)
		}
	} else {
		fmt.Println("no root")
	}
}

// 后序遍历链式， 打印顺序， 左子树，右子树，当前节点
func postOrderSearchList(btree *Node) {
	if (*btree).data != 0 {
		if (*btree).left != nil {
			postOrderSearchList((*btree).left)
		}

		if (*btree).right != nil {
			postOrderSearchList((*btree).right)
		}

		fmt.Println((*btree).data)
	} else {
		fmt.Println("no root")
	}
}


// 求高度
func getHight(btree *Node) int{
	var height int
	var leftHeight float64
	var rightHeight float64
	var leftPointer  = btree
	var rightPointer = btree
	for ; ; {
		if leftPointer.left != nil {
			leftPointer = leftPointer.left
			leftHeight += 1
		} else if rightPointer.right != nil {
			rightPointer = rightPointer.right
			rightHeight += 1
		} else {
			break
		}
	}

	height = int(math.Max(leftHeight, rightHeight))
	return height

}

func main() {
	fmt.Println("starting binary tree")
	var t1 Node
	generateListBinaryTree(&t1)


	//preOrderSearchList(&t1)
	//inOrderSearchList(&t1)
	postOrderSearchList(&t1)
	fmt.Println(t1)

	height := getHight(&t1)
	fmt.Println(height)
}
