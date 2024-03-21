package main

func ZeroToEnd() []int {
	
	arr := []int{1, 2, 0, 4, 3, 0, 5, 0}
	valuestore := 0

	for _, value := range arr {
		if value != 0 {
			arr[valuestore] = value
			valuestore++
		}
	}

	for ; valuestore < len(arr); valuestore++ {
		arr[valuestore] = 0
	}
	return arr
}
