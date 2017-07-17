package main

import "fmt"

func filter(numbers []int, callback func(int) bool) []int {
	var xs []int
	for _, n := range numbers {
		if callback(n) {
			fmt.Println("callback(n): ==> ", n)
			xs = append(xs, n)
		}
	}
	return xs
}

func main() {
	xs := filter([]int{1, 2, 3, 4}, func(n int) bool {
		fmt.Println("callback() call")
		return n > 1
	})
	fmt.Println("n > 1 ==> ", xs) // [2 3 4]
}
