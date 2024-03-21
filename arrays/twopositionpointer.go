package main

import "sort"

func twoPositionPointer(arr []int) {

	sort.Ints(arr)
	min := 0
	max := len(arr) - 1

	temp := make([]int, len(arr))

	for i := 0; i < len(arr); i++ {
		if i%2 == 0 {
			temp[i] = arr[max]
			max--
		} else {
			temp[i] = arr[min]
			min++
		}
	}
	copy(arr, temp)

}
