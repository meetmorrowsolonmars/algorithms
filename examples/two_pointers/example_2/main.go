package main

import (
	"fmt"
	"sort"
	"strings"
)

// Дан список неотрицательных целых чисел, повторяющихся элементов в списке нет.
// Нужно преобразовать его в строку, сворачивая соседние по числовому ряду числа в диапазоны.
// Примеры:
//
// [1, 4, 5, 2, 3, 9, 8, 11, 0] = > "0-5,8-9,11"
// [1, 4, 3, 2] = > "1-4"
// [1, 4] => "1,4"
// [1, 2] = > "1-2"
// [1] = > "1"
// [] = > ""

func main() {
	fmt.Println(convert([]int{1, 4, 5, 2, 3, 9, 8, 11, 0}))
}

func convert(list []int) string {
	if len(list) == 0 {
		return ""
	}

	// отсортировали в порядке возрастания
	sort.Ints(list)

	builder := strings.Builder{}
	left, right := list[0], list[0]

	for i := 1; i < len(list); i++ {
		if right+1 == list[i] {
			right = list[i]
			continue
		}

		if left == right {
			builder.Write([]byte(fmt.Sprintf("%d,", left)))
		} else {
			builder.Write([]byte(fmt.Sprintf("%d-%d,", left, right)))
		}

		left, right = list[i], list[i]
	}

	if left == right {
		builder.Write([]byte(fmt.Sprintf("%d", left)))
	} else {
		builder.Write([]byte(fmt.Sprintf("%d-%d", left, right)))
	}

	return builder.String()
}
