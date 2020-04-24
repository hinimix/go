package main

import (
	"fmt"
	"math/rand"
)

// 生成长度固定的随机数组
func generateRandomSlice(sliceName *[]int, n int) {
	for i := 0; i < n; i++ {
		(*sliceName)[i] = rand.Intn(5)
	}
}

// 快速排序
func quickSort(sliceName *[]int) []int {
	var length = len(*sliceName)
	var tmpResult = make([]int, 0)

	if length >= 2 {
		// 随机选取索引位置
		var randomPosition = rand.Intn(length)
		// 比索引小的数组存放位置
		var tmpSmaller = make([]int, 0)
		// 比索引大的数组存放位置
		var tmpBigger = make([]int, 0)
		var tmpMiddle = make([]int, 0)
		// 判断当前数组的每个元素，比随机索引小的放入smaller，比随机索引大的放入bigger
		for i := 0; i < length; i++ {
			if (*sliceName)[i] < (*sliceName)[randomPosition] {
				tmpSmaller = append(tmpSmaller, (*sliceName)[i])
			} else if (*sliceName)[i] > (*sliceName)[randomPosition] {
				tmpBigger = append(tmpBigger, (*sliceName)[i])
			} else {
				tmpMiddle = append(tmpMiddle, (*sliceName)[i])
			}
		}
		// 递归排序
		tmpSmaller = quickSort(&tmpSmaller)
		tmpBigger = quickSort(&tmpBigger)

		// 按顺序依次插入到tmpResult里返回，smaller, middle, bigger
		for i := 0; i < len(tmpSmaller); i++ {
			tmpResult = append(tmpResult, tmpSmaller[i])
		}
		for i := 0; i < len(tmpMiddle); i++ {
			tmpResult = append(tmpResult, tmpMiddle[i])
		}
		for i := 0; i < len(tmpBigger); i++ {
			tmpResult = append(tmpResult, tmpBigger[i])
		}
		// 如果长度只有1，直接返回
	} else if length == 1 {
		tmpResult = *sliceName
	}
	return tmpResult
}

// 二分查找，循环实现
func binarySearchLoop(sliceName *[]int, destinateValue int) int {
	//var result int
	var low = 0
	var high = len(*sliceName) - 1

	for ; ; {
		var mid = (low + high) / 2
		if low >= high {
			return mid
		} else if (*sliceName)[mid] > destinateValue {
			high = mid - 1
		} else if (*sliceName)[mid] < destinateValue {
			low = mid + 1
		} else if (*sliceName)[mid] == destinateValue {
			return mid
		}
	}
}

// 二分查找，递归实现
func binarySearchRecursive(sliceName *[]int, destinateValue int) int {
	var length = len(*sliceName)
	var low = 0
	var high = length - 1
	return binarySearchRecursiveStep(sliceName, low, high, destinateValue)
}

func binarySearchRecursiveStep(sliceName *[]int, low int, high int, destinateValue int) int {
	// 中间点
	var mid = (low + high) / 2
	//结果值
	var result = 65535

	if low >= high {
		result = mid
	} else if (*sliceName)[mid] > destinateValue {
		high = mid - 1
		result = binarySearchRecursiveStep(sliceName, low, high, destinateValue)

	} else if (*sliceName)[mid] < destinateValue {
		low = mid + 1
		result = binarySearchRecursiveStep(sliceName, low, high, destinateValue)
	} else if (*sliceName)[mid] == destinateValue {
		result = mid
	}
	return result
}

// 二分查找有重复值的第一个值
func binarySearchMany(sliceName *[]int, destinateValue int) int {
	var length = len(*sliceName)
	var low = 0
	var high = length - 1

	for ; ; {
		var mid = (low + high) / 2
		if low >= high {
			return mid
		} else if (*sliceName)[mid] > destinateValue {
			high = mid - 1
		} else if (*sliceName)[mid] < destinateValue {
			low = mid + 1
		} else if (*sliceName)[mid] == destinateValue && (*sliceName)[mid-1] != destinateValue {
			// 满足这个条件说明这是多个相同值的第一个节点
			return mid
		} else {
			// 否则缩小搜索范围继续搜索，小范围里也有目标值
			high = mid - 1
		}
	}
}

func main() {
	var n = 9
	var randomSlice1 = make([]int, n)
	var randomSlice2 = make([]int, 44)
	// 准备好排序的数组，以便二分查找
	generateRandomSlice(&randomSlice1, n)
	generateRandomSlice(&randomSlice2, n)
	randomSlice1 = quickSort(&randomSlice1)
	fmt.Println(randomSlice1)
	res1 := binarySearchLoop(&randomSlice1, 40)
	fmt.Println(res1)
	res2 := binarySearchRecursive(&randomSlice1, 40)
	fmt.Println(res2)

	randomSlice2 = quickSort(&randomSlice2)
	fmt.Println(randomSlice2)
	fmt.Println(binarySearchMany(&randomSlice2, 4))
}
