package loops

import "fmt"

func Loop_till_condition() {
	var i int
	for i < 5 {
		fmt.Println(i)
		i++
		if i == 4 {
			break
		}
		if i == 3 {
			continue
		}
		fmt.Println("continuing...")
	}
}

func Loop_till_condition_with_post_clause() {
	// var i int
	for i := 0; i < 5; i++ {
		print(i)
	}
}

func Loop_infinite() {
	var i int
	for {
		if i == 5 {
			break
		}
		i++
		fmt.Println(i)
	}
}

func Loop_over_collections() {
	slice := []int{1, 2, 3}
	// for i := 0; i < len(slice); i++ {
	// 	fmt.Println(slice[i])
	// }
	fmt.Println("looping over slice")
	for i, v := range slice {
		fmt.Println(i, v)
	}

	fmt.Println("looping over map")
	wellKnownPorts := map[string]int{"http": 80, "https": 443}
	for k, v := range wellKnownPorts {
		fmt.Println(k, v)
	}
	fmt.Println("looping over map: print only keys")
	for k := range wellKnownPorts {
		fmt.Println(k)
	}

	fmt.Println("looping over map: print only values")
	for _, v := range wellKnownPorts {
		fmt.Println(v)
	}
}
