package main

import (
	"fmt"
)

type slice []int

// func to merge two sorted array
func mergeSortedArrays(a, b slice) (rtn slice) {

	g := len(a)
	h := len(b)

	fmt.Println("array 1", a, "Length", g)
	fmt.Println("array 2", b, "Length", h)

	i := 0
	j := 0
	o := []int{}
	// work until length of whichever array is reached
	for j < h && i < g {
		if b[j] > a[i] {
			o = append(o, a[i])
			i++
		} else {
			if b[j] == a[i] {
				o = append(o, a[i], b[j])
				j++
				i++
			} else {
				if b[j] < a[i] {
					o = append(o, b[j])
					j++
				}
			}
		}
	}

	// append whatever remains of either array to output array
	for i < g {
		o = append(o, a[i])
		i++
	}

	for j < h {
		o = append(o, b[j])
		j++
	}
	return o
}

func main() {

	x := []int{1, 2, 3, 4, 6, 7, 8, 9, 10, 23, 40, 50, 50, 55, 60, 80, 90, 110, 120, 125}
	y := []int{4, 5, 7, 10, 11, 35, 39, 40, 45, 46, 49, 70, 75}
	p := mergeSortedArrays(x, y)
	fmt.Println("merged", p, "Length", len(p))
	q := mergeSortedArrays(y, x)
	fmt.Println("merged", q, "Length", len(q))
}
