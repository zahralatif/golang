package main

import "fmt"

func minMax(numbers []int) (min int, max int) {
	if len(numbers) == 0 {
		return 0, 0
	}
	min, max = numbers[0], numbers[0]
	for _, n := range numbers[1:] {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return
}

func main() {
	nums := []int{7, 21, 10, 14, 16}
	min, max := minMax(nums)
	fmt.Printf("Min: %d, Max: %d\n", min, max)
}
