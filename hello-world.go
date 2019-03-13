package main

import "fmt"

// import "math"
// import "time"
func smth(a int, b int) int {
	return a + b
}
func vals() (float32, float32) {
	return 26.5, 5.5
}
func sums(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
func main() {
	s := make([]string, 3)
	c := make([]string, len(s))
	s = []string{"a", "s", "d"}
	copy(c, s)
	fmt.Println("cpy:", c)
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	//range on arrays and slices provides both the index and value for each entry. Above we didn’t need the index, so we ignored it with the blank identifier _. Sometimes we actually want the indexes though.

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
	for j, num := range m {
		if num == 7 {
			fmt.Println("Found needed number with key: ", j)
		}
	}
	functCall := smth(2, 4)
	fmt.Println(functCall)
	a, b := vals()
	fmt.Println(a, b, "And sum is", a+b, "!")

	sums(1, 2)
	sums(1, 2, 3)

	nextInt := intSeq()
	//See the effect of the closure by calling nextInt a few times.

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newIntsss := intSeq()
	fmt.Println(newIntsss())
	fmt.Println(fact(5))
	//To confirm that the state is unique to that particular function, create and test a new one.

	newInts := intSeq()
	fmt.Println(newInts())

}
