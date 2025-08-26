package main

import "fmt"

func intersection(a, b []int) []int {
	seen := make(map[int]struct{}, len(a))
	for _, x := range a {
		seen[x] = struct{}{}
	}
	res := make([]int, 0)
	for _, y := range b {
		if _, ok := seen[y]; ok {
			res = append(res, y)
			delete(seen, y)
		}
	}
	return res
}

func main() {
	a := []int{1, 2, 3}
	b := []int{2, 3, 4}
	fmt.Println(intersection(a, b))
}
