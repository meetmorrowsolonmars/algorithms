package main

import "fmt"

func main() {
	testCases := []string{
		"()",
		"()[]{}",
		"(]",
		"([)]",
		"{[]}",
		"))",
	}

	for _, test := range testCases {
		fmt.Printf("Input: %s, Is valid: %v\n", test, isValid(test))
	}
}

// Функция для проверки корректности скобочной последовательности
func isValid(s string) bool {
	brackets := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := []rune{}

	for _, char := range s {
		if open, found := brackets[char]; found {
			if len(stack) == 0 || stack[len(stack)-1] != open {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, char)
		}
	}

	// Если стек пустой, последовательность корректна
	return len(stack) == 0
}
