package main

import "fmt"

// Найти все пары чисел в массиве, сумма которых равна заданному числу.
// Для входного массива [1, 5, 7, -1, 5] и целевого числа 6, программа должна вывести следующие пары:
// (1, 5)
// (7, -1)
// (1, 5) (опять же, так как есть еще один 5)

func main() {
	nums := []int{1, 5, 7, -1, 5}
	target := 6

	pairs := findPairs(nums, target)

	fmt.Println("Numbers:", nums)
	fmt.Println("Target:", target)
	fmt.Println("Pairs:", pairs)
}

// Функция для поиска всех пар чисел, сумма которых равна заданному числу target
func findPairs(nums []int, target int) [][2]int {
	// Создаем карту для хранения чисел, которые мы уже видели
	seen := make(map[int]struct{})
	// Слайс для хранения найденных пар
	pairs := [][2]int{}

	// Проходим по каждому числу в массиве
	for _, num := range nums {
		// Вычисляем комплимент текущего числа
		sub := target - num
		// Если разница чисел есть в карте, добавляем пару в результат
		if _, exists := seen[sub]; exists {
			pairs = append(pairs, [2]int{sub, num})
		}
		// Добавляем текущее число в карту
		seen[num] = struct{}{}
	}

	return pairs
}
