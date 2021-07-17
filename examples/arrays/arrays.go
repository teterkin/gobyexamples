package main

import "fmt"

func main() {

	var a [5]int
	fmt.Println("emp: ", a)

	a[4] = 100
	fmt.Println("set: ", a)
	fmt.Println("get: ", a[4])
	fmt.Println("len: ", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl: ", b)

	fmt.Println("---")

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		if i == 0 {
			fmt.Println("     j=0 j=1 j=2")
			fmt.Println("     ---|---|---")
		}
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
			if j == 0 {
				fmt.Printf("i=%d |", i)
			}
			fmt.Printf(" %d |", i+j)
		}
		fmt.Println()
		fmt.Println("     ---|---|---")
	}
	fmt.Println()
	fmt.Println("2d: ", twoD)

}
