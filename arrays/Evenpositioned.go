package main

import "fmt"

func EvenIsGreaterOdd(even []int) {

	if len(even) < 1 {
		return
	}

	for i := 0; i < len(even)-1; i += 2 {
		FirstElement := even[i]
		SecondElement := even[i+1]
		if even[i]%2 != 0 {
			if FirstElement > SecondElement {
				even[i], even[i+1] = FirstElement, SecondElement
				// fmt.Println(FirstElement,SecondElement)
			}
		} else if even[i]%2 == 0 {

			if FirstElement < SecondElement {
				even[i], even[i+1] = SecondElement, FirstElement
				// fmt.Println(FirstElement,SecondElement)
			}
		} else {
			fmt.Println("the given elemnts are negative numbers")
		}
	}

}

// func EvenIsGreaterOdd() {
//     even := make([]int, 10)
//     for i := 0; i < len(even)-1; i += 2 {
//         // Check if both elements are even or both are odd
//         if (even[i]%2 == 0 && even[i+1]%2 == 0) || (even[i]%2 != 0 && even[i+1]%2 != 0) {
//             // Swap if the first element is greater than the second
//             if even[i] > even[i+1] {
//                 even[i], even[i+1] = even[i+1], even[i]
//             }
//         }
//     }

//     fmt.Println(even)
// }
