package main

import "fmt"

func QuickSort(arr []int) []int{
	if len(arr) <= 1 {
		return arr
	}
	splitdata := arr[0] //取基数
	low := make([]int, 0, 0) //存储比基数小的数
	high := make([]int, 0 ,0) //存储比基数大的数
	mid := make([]int, 0, 0) //存储与基数相等的数

	mid = append(mid, splitdata) //将基数放到中间
	// 循环下标从1开始
	for i := 1; i < len(arr); i++ {
		if arr[i] < splitdata {
			low = append(low, arr[i]) //比基数小的放到左边
		}else if arr[i] > splitdata {
			high = append(high, arr[i]) //比基数大的放到右边
		}else {
			mid = append(mid, arr[i])
		}
	}
	//递归
	low = QuickSort(low)
	high = QuickSort(high)
	result := append(append(low, mid...), high...)
	return result
}

func main(){
	arr := []int{8, 9, 3, 6, 1, 7, 2, 15, 10, 11, 12, 4, 20, 18}
	fmt.Println(QuickSort(arr))
}
