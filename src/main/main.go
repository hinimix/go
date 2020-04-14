package main

import (
	"fmt"
)

type LIST struct {
	Data int
	Next *LIST
}

// 遍历链表
func showList(l *LIST) {
	for l != nil {
		fmt.Println(l.Data)
		//fmt.Println(reflect.TypeOf(l.Next))
		//fmt.Println(reflect.TypeOf(l.Data))
		//fmt.Println(reflect.TypeOf(l))
		l = l.Next
	}
}

// 插入链表尾部
func appendList(l *LIST, data LIST) {
	var tail *LIST
	for l.Next != nil {
		l = l.Next
	}
	tail = l
	tail.Next = &data
	data.Next = nil
}

// 插入链表头部
func insertList(l *LIST, data LIST) {
	var node1 = l.Next
	l.Next = &data
	data.Next = node1
}

// 生成一个链表
func generalList(name *LIST, l int) {
	for i := 0; i <= l; i++ {
		var tmpList LIST
		tmpList.Data = i
		appendList(name, tmpList)
	}
}

// 单链表反转
func reverceList(l *LIST) LIST {
	var result LIST
	// 判断当前是不是链表最后一个节点
	for l.Next != nil {
		l = l.Next
		insertList(&result, *l)
	}
	//insertList(&result, *l)
	return result
}

// 链表中环的检测: 判断所有链表节点的next有没有相同的，有就是有环路
func loopCheckList(l *LIST) bool {
	var result = false
	nextList := make([]*LIST, 5)

	var i = 0
	for l != nil {
		for j := 0; j <= i; j++ {
			if nextList[j] == l {
				//result=true
				return true
			}
		}
		nextList[i] = l
		l = l.Next
		i++
	}
	return result
}

// 制作一个环路的链表
func generateLoop() LIST {
	var t1, t2, t3 LIST
	t1.Next = &t2
	t2.Next = &t1
	t3.Next = &t2
	return t1
}

// 两个有序的链表合并
func mergeList(l1 *LIST, l2 *LIST) LIST {
	//var result LIST
	for l2 != nil {
		appendList(l1, *l2)
		l2 = l2.Next
	}
	return *l1
}

// 链表长度
func getLengthList(l *LIST) int {
	var count = 0
	for l.Next != nil {
		l = l.Next
		count += 1
	}
	return count
}

// 删除链表倒数第 n 个结点
func deleteNNode(l *LIST, n int) bool {
	// 执行结果
	var result = false
	// 链表长度
	var length = 0
	// 要删除的前一个结点
	var prevNode = l
	// 要删除的后一个结点
	var pastNode *LIST
	var countNode = l
	var realDeleteCount int

	// 计算链表长度

	length = getLengthList(countNode)

	// 判断正序倒序
	if n >= 0 {
		realDeleteCount = n
	} else {
		realDeleteCount = length + n
	}

	// 寻找要删除的前一个节点
	for i := 0; i < realDeleteCount; i++ {
		if prevNode.Next != nil {
			prevNode = prevNode.Next
		} else {
			fmt.Println("草")
		}

	}
	// 如果是-1，直接把前一个结点下一条赋值为空
	if n == -1 {
		prevNode.Next = nil
		result = true
	} else {
		pastNode = prevNode.Next
		pastNode = pastNode.Next
		prevNode.Next = pastNode
		result = true
	}
	return result
}

// 求链表的中间结点
func middleNodeList(l *LIST) LIST {
	var length = getLengthList(l)
	var middleNode = length/2 + length%2
	for i:=0;i<=middleNode;i++{
		l = l.Next
	}
	return *l
}

func main() {
	var t1 LIST
	t1.Data = -1
	generalList(&t1, 10)
	//insertList(&t1, LIST{5000, nil})
	appendList(&t1, LIST{5000, nil})
	//var t2 = reverceList(&t1)
	//showList(&t1)
	//showList(&t2)
	//var t3 = generateLoop()
	//var t3result = loopCheckList(&t3)
	//fmt.Println(t3result)
	//var t4 = mergeList(&t1, &t2)
	//showList(&t1)
	//var t5 = deleteNNode(&t1, -1)
	//fmt.Println(t5)
	showList(&t1)
	var t6 = middleNodeList(&t1)
	fmt.Println(t6.Data)
}
