package main

import "fmt"

/*
  首先建一个堆，然后调整堆，调整过程是将节点和子节点进行比较，将其中最大的值变为父节点，
  递归调整调整次数lgn,最后将根节点和尾节点交换再n次 调整O(nlgn).

  步骤：
    创建最大堆或者最小堆（我是最小堆）
    调整堆
    交换首尾节点(为了维持一个完全二叉树才要进行收尾交换)
*/

func main() {
	arr := []int{8, 9, 3, 6, 1, 7, 2, 15, 10, 11, 12, 4, 20, 18}
	fmt.Println(HeapSort(arr))
}

func HeapSortMax(arr []int, length int) []int {
	if length <= 1 {
		return arr
	}
	depth := length/2 - 1 //二叉树深度
	for i := depth; i >= 0; i-- {
		topMax := i //假定最大位置为i
		leftChild := 2*i + 1
		rightChild := 2*i + 2
		if (leftChild <= length-1) && (arr[leftChild] > arr[topMax]) { //大的作为父节点
			topMax = leftChild
		}
		if (rightChild <= length-1) && (arr[rightChild] > arr[topMax]) { //节点作为父节点
			topMax = rightChild
		}
		if topMax != i { //说明下标发生了交换
			arr[i], arr[topMax] = arr[topMax], arr[i] //值也要进行交换
		}
	}
	return arr
}

func HeapSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		lastlen := length - i
		HeapSortMax(arr, lastlen)
		if i < length {
			arr[0], arr[lastlen-1] = arr[lastlen-1], arr[0] //交换首尾节点
		}
	}
	return arr
}
