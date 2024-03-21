package main

func SecondLargestNumber(sli2 []int) int {
	if len(sli2) < 2 {
		return 0
	}
	largest := sli2[0]
	secondlargest := sli2[1]
	if largest < secondlargest {
		largest, secondlargest = secondlargest, largest
	}
	for i := 2; i < len(sli2); i++ {
		if sli2[i] > largest {
			secondlargest = largest
			largest = sli2[i]
		} else if sli2[i] > secondlargest && sli2[i] != largest {
			secondlargest = sli2[i]
		}
	}
	return secondlargest
}
