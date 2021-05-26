package main

import (
	"fmt"
)

/*
slice := sliceHeader{
    Length:        5,
    Capacity:      5,
    ZerothElement: &array[1],
}
*/

// len, swap, less

// Sort(array)

// {0,1,2,3,4}
type demoInt int

type demoFunc func(int)

func main() {
	// var array [10]int
	// array[0] = 50
	// fmt.Printf("address in main %p \n", &array[0])
	// var arr [10]demoInt
	// arr[0] = 1
	// sliceFromArray()
	// slice := []int{1, 2, 3, 4, 5}
	// fmt.Printf("%v", cap(slice[:3]))
}

func modifyArray(array *[10]int) {
	fmt.Printf("address in func %p \n", &array[0])
	array[0] = 10
}

func typesAllowedInArrays() {
	var array1 [10]demoInt
	array1[0] = 1
	var array2 [10]demoFunc
	// var array2 [10]func(int)
	array2[0] = func(i int) { fmt.Println(i) }
	array2[0](5)
}

func sliceFromArray() {

	var array [6]int = [6]int{0, 1, 2, 3, 4, 5}
	slice := array[1:]
	// 0 1 2 3 4
	// 1 2 3 4 5
	s1 := slice[:3] // [4,5]
	fmt.Println(cap(s1))
	slice[0] = 100
	fmt.Println(array)
	fmt.Println(slice)

	fmt.Println("capacity of slice: ", cap(slice))
	fmt.Println("len of slice:", len(slice))
	fmt.Printf("address for array %p \n", &array[1])
	fmt.Printf("address for slice %p \n", &slice[0])

	slice = append(slice, 1)

	fmt.Println("capacity of slice after append: ", cap(slice))
	fmt.Println("len of slice after append:", len(slice))
	fmt.Printf("address for array after append %p \n", &array[1])
	fmt.Printf("address for slice after append %p \n", &slice[0])
}

func useOfMake() {
	//make([]T, len, cap)
	s := make([]int, 0, 5)
	fmt.Println(s)
}

func literals() {
	arr := [5]int{1, 2, 3, 4, 5}
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	fmt.Println(slice)
}

func slicing() {
	// slice := []int{1, 2, 3, 4, 5}
	// s1:=[:]
	// s2:=[:2]
	// s3:=[3:]
}

func copying() {
	a := []int{1, 2}
	b := make([]int, len(a))
	copy(b, a)
}

func sliceAppend() {
	a := []int{1, 2}
	b := []int{3, 4, 5}
	a = append(a, b...)
}
