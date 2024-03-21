package main

//find the largest three numbers in an array
import (
	"fmt"
)

func findlargestthreenumbers(sli []int) (int, int, int) {
	if len(sli) < 3 {
		return 0, 0, 0
	}
	first, second, third := sli[0], sli[1], sli[2]

	if first < second {
		first, second = second, first
	}
	if first < third {
		first, third = third, first
	}
	if second < third {
		second, third = third, second
	}

	for i := 3; i < len(sli); i++ {
		if sli[i] > first {
			third = second
			second = first
			first = sli[i]
		} else if sli[i] > second {
			third = second
			second = sli[i]
		} else if sli[i] > third {
			third = sli[i]
		}
	}
	return first, second, third
}

// func rearrangingElements(){
// 	sli3:=[]int{1,40,2,30,50,4,6,90,100}
//     firstelement,secondelement:=sli3[0],sli3[1]

// 	for i,value := range sli3{
// 		if firstelement>secondelement{
// 			firstelement,secondelement=secondelement,firstelement
// 		}
// 	}
// }

func main() {

	//	sli :=[]int{10,20,30,40,60,70,100,200}
	//	first,second,third := findlargestthreenumbers(sli)
	//	fmt.Println("Largest three elements ",first,second,third)
	//	sli2 := []int{10, 30, 50, 60, 80, 300}
	//	secondnumber := SecondLargestNumber(sli2)
	//
	// fmt.Println("Second largest number:",secondnumber )
	//
		zero := ZeroToEnd()
	
	fmt.Println("Zero elements to the end", zero)
	// sli := []int{1, 2, 3, 4, 5, 6, 7, 8}
	// EvenIsGreaterOdd(sli)
	// fmt.Println(sli)
	// arr := []int{1, 2, 3, 4, 5, 6, 7}
	// twoPositionPointer(arr)

	// fmt.Println(arr)

}
