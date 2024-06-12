package main

import "fmt"

// Для заданного отсортированного в возрастающем порядке массива A целых чисел длины n и целого числа X определить,
// существует ли в нем два различных элемента A[i] и A[j], такие что A[i] + A[j] = X.

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 15, 20, 30, 40, 50, 55, 60, 70}
	x := 60
	found := false

	l, r := 0, len(a)-1
	for l < r {
		sum := a[l] + a[r]
		if a[l]+a[r] == x {
			found = true
			break
		}

		if sum > x {
			r--
		} else {
			l++
		}
	}

	if found {
		fmt.Printf("a[%d] + a[%d] = %d + %d = %d\n", l, r, a[l], a[r], a[l]+a[r])
	} else {
		fmt.Println("not found")
	}
}
