package main

import (
	"fmt"
	"math/rand"
)

// 生成长度固定的随机数组
func generateRandomSlice(sliceName *[]int, n int) {
	for i := 0; i < n; i++ {
		(*sliceName)[i] = rand.Intn(10)
	}
}

// 冒泡排序
func bubbleSort(sliceName *[]int) {
	for x := 0; x < len(*sliceName); x++ {
		for y := x; y < len(*sliceName); y++ {
			//相邻元素比较，如果前面比后面大，交换位置
			if (*sliceName)[x] > (*sliceName)[y] {
				var tmp = (*sliceName)[x]
				(*sliceName)[x] = (*sliceName)[y]
				(*sliceName)[y] = tmp
			}
		}
	}
}

// 插入排序
func insertSort(sliceName *[]int) {
	// 如果数组<=1，不用排序
	if len(*sliceName) <= 1 {
		return
	}
	//x是未排序分组
	for x := 1; x < len(*sliceName); x++ {
		value := (*sliceName)[x]
		//y是已排序分组
		for y := 0; y < x; y++ {
			if (*sliceName)[y] > value {
				(*sliceName)[x] = (*sliceName)[y]
				(*sliceName)[y] = value
				value = (*sliceName)[x]
			}
		}
	}
}

// 选择排序
func selectSort(sliceName *[]int) {
	if len(*sliceName) <= 1 {
		return
	}
	// 从0开始遍历搜索最小的赋值到数组前面
	for x := 0; x < len(*sliceName); x++ {
		min_value := (*sliceName)[x]
		for y := 1; y < len(*sliceName); y++ {
			if (*sliceName)[x] > (*sliceName)[y] {
				(*sliceName)[x] = (*sliceName)[y]
				(*sliceName)[y] = min_value
				min_value = (*sliceName)[y]
			}
		}
	}
}

// 归并排序
func mergeSort(sliceName *[]int) []int {
	length := len(*sliceName)
	tmpHead := make([]int, length/2)
	tmpTail := make([]int, length-length/2)
	tmpResult := make([]int, 0)

	if length > 2 {
		// 把一个数组拆分成两个
		for i := 0; i < length; i++ {
			if i < length/2 {
				tmpHead[i] = (*sliceName)[i]
			} else {
				tmpTail[i-length/2] = (*sliceName)[i]
			}
		}
		// 递归进去排序
		tmpHead = mergeSort(&tmpHead)
		tmpTail = mergeSort(&tmpTail)

		// 两个坐标来同时控制两个数组
		var x = 0
		var y = 0
		for i := 0; i < length; i++ {
			if x >= len(tmpHead) {
				// 如果前面数组都空了，就插入后面数组
				tmpResult = append(tmpResult, tmpTail[y])
				y++
			} else if y >= len(tmpTail) {
				// 如果后面数组都空了，就插入前面数组
				tmpResult = append(tmpResult, tmpHead[x])
				x++
			} else if tmpHead[x] >= tmpTail[y] {
				// 比较大小，小的进临时数组
				tmpResult = append(tmpResult, tmpTail[y])
				y++
			} else if tmpHead[x] < tmpTail[y] {
				tmpResult = append(tmpResult, tmpHead[x])
				x++
			}
		}

	} else if length < 2 {
		tmpResult = (*sliceName)
	} else {
		if (*sliceName)[0] > (*sliceName)[1] {
			tmpResult = append(tmpResult, (*sliceName)[1])
			tmpResult = append(tmpResult, (*sliceName)[0])
		} else {
			tmpResult = append(tmpResult, (*sliceName)[0])
			tmpResult = append(tmpResult, (*sliceName)[1])
		}
	}
	return tmpResult
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

// 计数排序
func counterSort(sliceName *[]int) []int {
	var bucketLength = 10
	var bucket = make([]int, bucketLength)
	var tmpResult = make([]int, len(*sliceName))

	// 判断每个bucket的数值
	for i := 0; i < len(*sliceName); i++ {
		for y := 0; y < bucketLength; y++ {
			if (*sliceName)[i] == y {
				bucket[y]++
			}
		}
	}

	// 求和
	sum := 0
	for i := 0; i < bucketLength; i++ {
		sum += bucket[i]
	}
	index := sum - 1
	//把符合数组的值插入到tmpResult
	for  ;index>=0;{
		for j:= bucket[bucketLength-1]; j>0; j-- {
			tmpResult[index] = bucketLength-1
			index--
		}
		bucketLength--
	}
	//for i:=len(*sliceName)-1;i>=0;i--{
	//	var index = bucket[(*sliceName)[i]]-1
	//	tmpResult[index] = (*sliceName)[i]
	//	bucket[(*sliceName)[i]]--
	//}


	return tmpResult
}

func main() {
	var n = 22
	var randomSlice1 = make([]int, n)
	var randomSlice2 = make([]int, n)
	var randomSlice3 = make([]int, n)
	var randomSlice4 = make([]int, n)
	var randomSlice5 = make([]int, n)
	var randomSlice6 = make([]int, n)
	generateRandomSlice(&randomSlice1, n)
	fmt.Println(randomSlice1)
	bubbleSort(&randomSlice1)
	fmt.Println("bubble sort: ", randomSlice1)

	generateRandomSlice(&randomSlice2, n)
	fmt.Println(randomSlice2)
	insertSort(&randomSlice2)
	fmt.Println("insert sort: ", randomSlice2)

	generateRandomSlice(&randomSlice3, n)
	fmt.Println(randomSlice3)
	insertSort(&randomSlice3)
	fmt.Println("select sort: ", randomSlice3)

	generateRandomSlice(&randomSlice4, n)
	fmt.Println(randomSlice4)
	randomSlice4 = mergeSort(&randomSlice4)
	fmt.Println("merge sort: ", randomSlice4)

	generateRandomSlice(&randomSlice5, n)
	fmt.Println(randomSlice5)
	randomSlice5 = quickSort(&randomSlice5)
	fmt.Println("quicksort: ", randomSlice5)

	generateRandomSlice(&randomSlice6, n)
	fmt.Println(randomSlice6)
	randomSlice6 = counterSort(&randomSlice6)
	fmt.Println("countersort: ", randomSlice6)

}
