package main

import "fmt"

/*
    算法描述：在一组有序数组中，将数组一分为二，将要查询的元素和分割点进行比较，分为三种情况
    1.相等直接返回
    2.元素大于分割点，在分割点右侧继续查找
    3.元素小于分割点，在分割点左侧继续查找
    时间复杂： O(lgn).
    要求:必须是有序的数组，并能支持随机访问
 */

func BinarySearch(arr []int, data int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] > data {
			high = mid - 1
		}else if arr[mid] < data {
			low = mid + 1
		}else {
			return mid
		}
	}
	return -1
}

func main()  {
	arr := make([]int, 1024 * 1024, 1024 * 1024)
	for i := 0; i < 1024 * 1024; i++ {
		arr[i] = i + 1
	}
	index := BinarySearch(arr, 1024)
	if index != -1 {
		fmt.Printf("index = %d\n", index)
	}else {
		fmt.Println("未找到数据")
	}
}
