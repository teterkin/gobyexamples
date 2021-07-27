package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	} else if n == 1 {
		fmt.Printf("%v = ", n)
		return n * fact(n-1)
	} else {
		fmt.Printf("%v * ", n)
		return n * fact(n-1)
	}
}

func main() {

	// 7 * 6 * 5 * 4 * 3 * 2 * 1 = 5040

	fmt.Println(fact(7))

}
