package main

import "fmt"

func main() {
	arr := []int {1,2,3,4,5,6}
	var result = make([]int,len(arr))
	mergeSort(arr,0,len(arr) - 1,result)
	fmt.Println(result)
}

func mergeSort(arr []int ,top int,rear int,result []int)  {
	if top < rear {
		mid := (top + rear) / 2
		mergeSort(arr,top,mid,result)
		mergeSort(arr,mid + 1,rear,result)
		mergeArray(arr,top,mid,rear,result)
	}
}

func mergeArray(arr []int, top int, mid int, rear int, result []int) {
	i := top
	j := mid + 1
	index := 0
	for i <= mid && j <= rear  {
		if arr[i] <= arr[j] {
			result[index] = arr[i]
			index++
			i++
		} else {
			result[index] = arr[j]
			index++
			j++
		}
	}
	
	for i <= mid {
		result[index] = arr[i]
		index++
		i++
	}
	
	for j <= rear {
		result[index] = arr[j]
		index++
		j++
	}
	
	// 保证切片 arr 有序
	for i = 0; i < index; i++ {
		arr[top+i] = result[i]
	}
}
