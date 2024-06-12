package main

import "fmt"

func main() {
	nums1 := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k1 := 3
	fmt.Printf("Input: %v, k: %d\nOutput: %v\n", nums1, k1, maxSlidingWindow(nums1, k1))

	nums2 := []int{1}
	k2 := 1
	fmt.Printf("Input: %v, k: %d\nOutput: %v\n", nums2, k2, maxSlidingWindow(nums2, k2))
}

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	var result []int
	// Используем как двустороннюю очередь.
	// Будем добавлять новые значения в конец списка, а удалять элементы сначала.
	deque := []int{}

	// nums1 := []int{1, 3, -1, -3, 5, 3, 6, 7}
	// k1 := 3
	// i = 0) deque = {1}
	// i = 1) deque = {3}
	// i = 2) deque = {3, -1}, result = {3}
	// i = 3) deque = {3, -1, -3}, result = {3, 3}
	// i = 4) deque = {5}, result = {3, 3, 5}
	// i = 5) deque = {5, 3}, result = {3, 3, 5, 5}
	// i = 6) deque = {6}, result = {3, 3, 5, 5, 6}
	// i = 7) deque = {7}, result = {3, 3, 5, 5, 6, 7}
	// result = {3, 3, 5, 5, 6, 7}

	for i := 0; i < len(nums); i++ {
		// Удаляем элементы, которые вышли за границу окна.
		// Элементы удаляются не каждую итерацию цикла, а только когда окно перестало их захватывать.
		if len(deque) > 0 && deque[0] < i-k+1 {
			deque = deque[1:]
		}

		// Удаляем элементы из конца deque, которые меньше текущего элемента.
		// Таким образом слева (в начале очереди) всегда будет самый большой элемент текущего окна.
		// Но так как окно двигается, то элементы, значение которых не больше значения последнего
		// элемента в очереди мы сохраняем.
		for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[i] {
			deque = deque[:len(deque)-1]
		}

		// Добавляем текущий элемент в deque.
		deque = append(deque, i)

		// Добавляем максимальный элемент текущего окна в результат.
		// Окно двигается каждую итерацию после этой отметки.
		// Поэтому пропустив k элементов мы дальше добавляем значения в результат на каждой итерации.
		if i >= k-1 {
			result = append(result, nums[deque[0]])
		}
	}

	return result
}
