package main

import (
	"fmt"
	"sort"
)

func main() {
	setA := []int{2, 1, 3, 5, 10}
	setB := []int{5, 4, 20, 10}

	fmt.Println("A: ", setA)
	fmt.Println("B: ", setB)

	IntersecSort(setA, setB)
	IntersecMap(setA, setB)
	IntersecForce(setA, setB)
}

func IntersecForce(setA []int, setB []int) {
	intersection := make([]int, 0)
	for _, a := range setA {
		for _, b := range setB {
			if a == b {
				intersection = append(intersection, b)
			}
		}
	}
	fmt.Println("Intersection force: ", intersection)
}

func IntersecMap(setA []int, setB []int) {
	a := make(map[int]struct{}, 0)
	for _, el := range setA {
		a[el] = struct{}{}
	}
	intersection := make([]int, 0)
	for _, el := range setB {
		if _, found := a[el]; found {
			intersection = append(intersection, el)
		}
	}
	fmt.Println("Intersection map", intersection)
}

func IntersecSort(setA []int, setB []int) {
	sort.Ints(setA)
	sort.Ints(setB)

	intersection := make([]int, 0)

	for i, j := 0, 0; i < len(setA) && j < len(setB); {
		if setA[i] == setB[j] {
			intersection = append(intersection, setB[j])
			i++
			j++
		} else if setA[i] < setB[j] {
			i++
		} else {
			j++
		}
	}
	fmt.Println("Intersection sort: ", intersection)
}
