package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}
	slice := arr[:]
	arr[1] = 42
	slice[2] = 27
	fmt.Println(arr, slice)

	slice_2 := []int{1, 2, 3}
	fmt.Println(slice_2)
	slice_2 = append(slice_2, 4)
	fmt.Println(slice_2)

	s3 := slice_2[1:]
	s4 := slice_2[:2]
	s5 := slice_2[1:2]

	fmt.Println(s3, s4, s5)
}
