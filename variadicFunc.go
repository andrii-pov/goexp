package main

import "fmt"

func sum(nums ...int) (total int) {
	total = 0
	for _, num := range nums {
		total += num
	}
	
	return
}

func main() {
	sum(1,3)

	nums := []int{1,2,44,3}
	fmt.Println(sum(nums...))
	
}
