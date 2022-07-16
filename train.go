package main

import "fmt"

func main() {
	var a, z int
	fmt.Scan(&a)
	var m []int
	min := 100
	count := 0
	for i := 0; i < a; i++ {
		fmt.Scan(&z)
		m = append(m, z)
	}
	for _, z := range m {
		if z < min {
			min = z
		}
	}
	for _, r := range m {
		if r == min {
			count += 1
		}
	}
	fmt.Println(count)
}
