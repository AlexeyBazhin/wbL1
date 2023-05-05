package main

import "fmt"

func main() {
	for key := range intersectionMap([]int{1,2,3,4,5,6}, []int{3,2,3,2, 5,5,5}) {
		fmt.Print(key, " ")
	}
	fmt.Println()
}

func intersectionMap(set1, set2 []int) map[int]struct{}{
	// создаем map для хранения элементов первого множества
	set1Map := make(map[int]struct{})
	for _, v := range set1 {
		set1Map[v] = struct{}{}
	}

	// находим пересечение второго множества с map первого множества
	res := make(map[int]struct{})
	for _, v := range set2 {
		if _, ok := set1Map[v]; ok {
			res[v] = struct{}{}
		}
	}

	return res
}
