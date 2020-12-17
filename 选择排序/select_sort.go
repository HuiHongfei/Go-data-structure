package main

import "fmt"
/*
   算法描述：从未排序数据中选择最大或者最小的值和当前值交换 O(n^2).
   算法步骤:
   1.选择一个数当最小值或者最大值，进行比较然后交换
   2.循环向后查进行
*/

//获取切片里面的最大值
func SelectMax(arr []int) int {
	length := len(arr)
	if length <= 1 {
		return arr[0]
	}
	max := arr[0]
	for i := 1; i < length; i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

//切片排序
func SelectSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	for i := 0; i < length; i++ {
		min := i
		// 每一轮找的都是arr里面最小数的下标
		for j := i + 1; j < length; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		// 一轮找完之后与arr中未排序的首元素交换
		if i != min {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}
	return arr
}

//选择排序
func main() {
	arr := []int{8, 9, 3, 6, 1, 7, 2, 15, 10, 11, 12, 4, 20, 18}
	max := SelectMax(arr)
	selectSort := SelectSort(arr)
	fmt.Println(max)
	fmt.Println(selectSort)
}
