package main

import (
	"fmt"
	"math/rand"
)

// 堆和堆排序
// 堆是完全二叉树，用数组合适
// 这里的堆都是大顶堆

type Heap struct {
	// 数组，装载完全二叉树
	data []int
	// 数组容量
	capacity int
	// 数组长度
	length int
}

// 堆化交换，自下而上
func (h Heap) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

// 获得父节点索引
func (h Heap) getParent(pos int) int {
	if pos == 0 {
		return 0
	} else if pos%2 == 1 {
		return (pos - 1) / 2
	} else if pos%2 == 0 {
		return (pos - 2) / 2
	}
	return 65535
}

// 堆化函数，迭代与父节点比对交换，直到根节点，自下而上
func (h Heap) heapifyUp(index int) {
	if index == 0 {
		return
	}
	parentIndex := h.getParent(index)

	if h.data[parentIndex] < h.data[index] {
		h.swap(parentIndex, index)
		h.heapifyUp(parentIndex)
	} else {
		// 不需要换位置
		return
	}

}

// 堆插入
func insertHeap(heap *Heap, value int, pos int) {
	// 位置大于数组长度，堆满了
	if pos > (*heap).capacity {
		fmt.Println("heap is full")
		return
	}

	// 新元素直接插入到结尾
	if (*heap).data[pos] == 0 {
		(*heap).data[pos] = value
	} else {
		(*heap).data[(*heap).length] = value
	}

	(*heap).heapifyUp((*heap).length)
	(*heap).length++

}

// 随机生成堆
func generateHeap(heap *Heap) {
	for i := 0; i < 50; i++ {
		//fmt.Println(i)
		randomValue := rand.Intn(100)
		insertHeap(heap, randomValue, 0)
	}
}

// 堆化，自上而下
func (h Heap) heapifyDown(index int) {
	if index*2+1 < h.length && h.data[index] < h.data[index*2+1] {
		h.swap(index, index*2+1)
		h.heapifyDown(index*2 + 1)
	} else if index*2+2 <= h.length && h.data[index] < h.data[index*2+2] {
		h.swap(index, index*2+2)
		h.heapifyDown(index*2 + 2)
	} else {
		return
	}
}

// 建堆
func (h Heap) createHeap(){
	for i := h.length/2; i >= 0; i-- {
		h.heapifyDown(i)
	}
}

// 堆排序
func heapSort(heap *Heap) {
	// 建堆, 自上而下堆化
	(*heap).createHeap()
	// 排序, 把大顶堆的最大值插入到最后节点，把次大值插入n-1的位置
	for i:=(*heap).length-1; i>=1; i-- {
		// 把堆首最大值交换到末尾
		(*heap).swap(0, i)
		//(*heap).data[0] = (*heap).data[(*heap).length]
		(*heap).length --
		(*heap).createHeap()
	}

}

func main() {
	//var hp = make([]int, 50)
	var hp Heap
	hp.capacity = 50
	hp.data = make([]int, hp.capacity)
	generateHeap(&hp)
	fmt.Println(hp.data)

	var data = make([]int, 15)
	for i := 0; i < 15; i++ {
		data[i] = rand.Intn(100)
	}
	fmt.Println(data)

	var hp1 Heap
	hp1.data = data
	hp1.capacity = 15
	hp1.length = 15

	heapSort(&hp1)
	fmt.Println(hp1.data)
}
